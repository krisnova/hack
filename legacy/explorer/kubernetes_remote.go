/*
Copyright © 2020-2021 Kris Nóva <kris@nivenly.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package explorer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kris-nova/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/deprecated/scheme"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type RemoteExplorerOptions struct {
	KubeconfigPath string
	Namespace      string
	Image          string
	Shell          string
	Name           string

	// ----------------------
	// [  Security Section  ]
	//
	MountProc           bool
	PrivilegeEscalation bool
	HostNetwork         bool
	HostIPC             bool
	HostPID             bool
	GroupID             int
	UserID              int

	//
	// ----------------------

}

type RemoteExplorer struct {
	ClientSet kubernetes.Clientset
	Config    rest.Config
	Options   *RemoteExplorerOptions
}

func NewRemoteExplorer(clientSet kubernetes.Clientset, config rest.Config, options *RemoteExplorerOptions) *RemoteExplorer {
	return &RemoteExplorer{
		ClientSet: clientSet,
		Config:    config,
		Options:   options,
	}
}

type LearnedPrivilege bool

type ProbedProfile struct {
	ClusterName               string
	Nodes                     []v1.Node
	AccessKubeSystemNamespace LearnedPrivilege
}

//
// Probe is what is used to determine what privileges we can
// run with.
//
func (e *RemoteExplorer) Probe() (*ProbedProfile, error) {

	logger.Always("Building [REMOTE] Profile")
	profile := &ProbedProfile{}

	// Nodes
	logger.Always("Probing Nodes...")
	nodes, err := e.ClientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	for _, node := range nodes.Items {
		profile.ClusterName = node.ClusterName
		addrs := ""
		for _, addr := range node.Status.Addresses {
			addrs = addrs + " " + addr.Address
		}
		logger.Success("[Node %s] %s", node.Name, addrs)
		profile.Nodes = append(profile.Nodes, node)
	}

	// Kube System Namespace
	logger.Always("Probing kube-system namespace")
	_, err = e.ClientSet.CoreV1().Namespaces().Get(context.Background(), "kube-system", metav1.GetOptions{})
	if err != nil {
		profile.AccessKubeSystemNamespace = false
		logger.Warning("[AccessKubeSystemNamespace] DENIED")
	} else {
		profile.AccessKubeSystemNamespace = true
		logger.Success("[AccessKubeSystemNamespace] GRANTED")
	}
	return profile, nil
}

func i64(i int) *int64 {
	pi := int64(i)
	return &pi
}

func b(i bool) *bool {
	return &i
}

//
// Attach will attach to a pod based on input from the user
//
func (e *RemoteExplorer) Attach(profile *ProbedProfile, namespace, image, shell string) error {
	//
	// Set up the attachment. Here we define the Pod and declare our pod configuration.
	//
	name := newName()
	if e.Options.Name != "" {
		name = e.Options.Name
	}
	logger.Always("Creating pod: %s", name)

	//
	// Let's set up our security context based on the user input
	//
	procMount := v1.DefaultProcMount // This will use the default container runtime /proc configuration
	if e.Options.MountProc {
		logger.Success("Mounting /proc from host.")
		procMount = v1.UnmaskedProcMount // This WILL mount /proc as it is on the host :)
	}

	containerName := newName()
	logger.Always("Container name: %s", name)
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		TypeMeta: metav1.TypeMeta{},
		Spec: v1.PodSpec{
			HostIPC:     e.Options.HostIPC,
			HostPID:     e.Options.HostPID,
			HostNetwork: e.Options.HostNetwork,
			SecurityContext: &v1.PodSecurityContext{
				RunAsGroup: i64(e.Options.GroupID), // GID 0
				RunAsUser:  i64(e.Options.UserID),  // UID 0
			},
			Containers: []v1.Container{
				v1.Container{
					Name:            containerName,
					Image:           image,
					ImagePullPolicy: v1.PullAlways,
					SecurityContext: &v1.SecurityContext{
						AllowPrivilegeEscalation: b(e.Options.PrivilegeEscalation), // Allow setns()
						Privileged:               b(e.Options.PrivilegeEscalation), // Access to the host
						RunAsGroup:               i64(e.Options.GroupID),           // GID 0
						RunAsUser:                i64(e.Options.UserID),            // UID 0
						ProcMount:                &procMount,                       // Defined above the /proc permissions
					},
				},
			},
		},
	}
	options := metav1.CreateOptions{}
	//
	// Create the Pod. This is where we start to mutate the cluster.
	// Make sure we defer() removing the Pod
	//
	pod, err := e.ClientSet.CoreV1().Pods(namespace).Create(context.Background(), pod, options)
	// Here we defer the pod deletion to the end of the function
	e.DeferDeletePod(name, namespace)
	defer func() {
		logger.Always("Deleting pod: %s", name)
		err := e.ClientSet.CoreV1().Pods(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
		if err != nil {
			logger.Warning("Error deleting pod: %v", err)
		}
	}()
	if err != nil {
		return err
	}
	//
	// Hang on Pod creation
	//
	{
		i := 1000 // Try for 3000 seconds
		for {
			if i == 0 {
				return fmt.Errorf("unable to attach to pod %s", name)
			}
			pod, err := e.ClientSet.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
			if err != nil {
				logger.Warning(err.Error())
				time.Sleep(time.Second * 3)
				i--
				continue
			}
			if pod.Status.Phase == "Running" {
				logger.Always("Pod %s: Running", name)
				break
			}
			logger.Always("Pod %s: %s", name, pod.Status.Phase)
			time.Sleep(time.Second * 3)
			i--
			continue
		}
	}
	{
	}

	logger.Always("Attaching to pod: %s", name)
	cmd := []string{
		shell,
	}
	request := e.ClientSet.CoreV1().RESTClient().Post().Resource("pods").Name(name).Namespace(namespace).SubResource("exec")
	option := &v1.PodExecOptions{
		Command: cmd,
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}
	request.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(&e.Config, "POST", request.URL())
	if err != nil {

		// If we hit an error, most likely the cluster has rejected
		// our request -- let's let the user know what happened
		pod, err2 := e.ClientSet.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
		if err2 != nil {
			return fmt.Errorf("Error executing pod %v: Error finding error: %v", err, err2)
		}
		jsonBytes, err2 := json.Marshal(pod)
		if err != nil {
			return fmt.Errorf("Error executing pod: %v Error JSON: %v", err, err2)
		}
		podStr := string(jsonBytes)
		logger.Debug(podStr)
		return fmt.Errorf("Error executing pod: %v \n\n %s", podStr)
	}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		return err
	}
	return nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func newName() string {
	rand.Seed(time.Now().UnixNano())
	n := 10
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
