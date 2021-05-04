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
package cmd

import (
	"os"
	"path/filepath"

	"github.com/kris-nova/hack/explorer"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// itCmd represents the it command
var itCmd = &cobra.Command{
	Use:   "it",
	Short: "Interactive TTY (EG: docker run -it). Run a shell in Kubernetes.",
	Long: `

This is where the magic happens. Here is where we open up a remote shell
in a Kubernetes cluster with as many privileges as possible.

`,
	Run: func(cmd *cobra.Command, args []string) {
		// Process arguments here
		home := homedir.HomeDir()
		if home == "" {
			logger.Critical("unable to parse $HOME")
			os.Exit(2)
		}
		workingPath := filepath.Join(home, ".kube", "config")
		itopt.KubeconfigPath = workingPath // Update the manipulated path
		err := InteractiveTTY(itopt)
		if err != nil {
			logger.Critical(err.Error())
			os.Exit(1)
		}
		logger.Always("Exit...")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(itCmd)
	itCmd.Flags().StringVarP(&itopt.KubeconfigPath, "kubeconfig", "k", "~/.kube/config", "The path on the local filesystem to read a Kube config file from.")
	itCmd.Flags().StringVarP(&itopt.Namespace, "namespace", "n", "default", "The namespace in Kubernetes to attach to.")
	itCmd.Flags().StringVarP(&itopt.Image, "image", "i", "krisnova/hack:latest", "The container image to run.")
	itCmd.Flags().StringVarP(&itopt.Shell, "shell", "s", "/bin/bash", "The path to the shell or command to run EX: /bin/bash")
	itCmd.Flags().StringVarP(&itopt.Name, "name", "N", "", "Name to use for the pod.")
	itCmd.Flags().IntVarP(&itopt.GroupID, "gid", "g", 0, "The group ID (GID) to run the container with.")
	itCmd.Flags().IntVarP(&itopt.GroupID, "uid", "u", 0, "The user ID (UID) to run the container with.")
	itCmd.Flags().BoolVarP(&itopt.PrivilegeEscalation, "privileged", "p", true, "Controls both the privileged and allowPrivilegedEscalation bools")
	itCmd.Flags().BoolVarP(&itopt.MountProc, "mount-proc", "m", true, "Controls the masking for proc. If enabled will attempt to mount /proc from the host.")
	itCmd.Flags().BoolVarP(&itopt.HostPID, "pid", "P", true, "Share the host pid namespace with the pod.")
	itCmd.Flags().BoolVarP(&itopt.HostNetwork, "network", "x", true, "Share the host network namespace with the pod.")
	itCmd.Flags().BoolVarP(&itopt.HostIPC, "ipc", "I", true, "Share the IPC network namespace with the pod.")

}

var itopt = &explorer.RemoteExplorerOptions{}

//
//
// InteractiveTTY is an opinionated way to enter a pod in a Kubernetes cluster.
//
func InteractiveTTY(itopt *explorer.RemoteExplorerOptions) error {
	logger.Always("Interactive TTY")
	explorer.HandleSignals()
	config, err := clientcmd.BuildConfigFromFlags("", itopt.KubeconfigPath)
	if err != nil {
		return err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	remoteEx := explorer.NewRemoteExplorer(*clientSet, *config, itopt)
	profile, err := remoteEx.Probe()
	if err != nil {
		return err
	}
	err = remoteEx.Attach(profile, itopt.Namespace, itopt.Image, itopt.Shell)
	if err != nil {
		return err
	}
	return nil
}
