package explorer

import (
	"context"
	"github.com/kris-nova/logger"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type RemoteExplorer struct {
	ClientSet kubernetes.Clientset
}

func NewRemoteExplorer (clientSet kubernetes.Clientset) *RemoteExplorer{
	return &RemoteExplorer{
		ClientSet: clientSet,
	}
}

type LearnedPrivilege bool

type ProbedProfile struct {
	ClusterName string
	Nodes []v12.Node
	AccessKubeSystemNamespace LearnedPrivilege
}

/**
 * Probe is what is used to determine what privileges we can
 * run with.
 */
func (e *RemoteExplorer) Probe() (*ProbedProfile, error) {

	logger.Always("Building [REMOTE] Profile")
	profile := &ProbedProfile{}

	// Nodes
	logger.Always("Probing Nodes...")
	nodes, err := e.ClientSet.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
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
	_, err = e.ClientSet.CoreV1().Namespaces().Get(context.Background(), "kube-system", v1.GetOptions{})
	if err != nil {
		profile.AccessKubeSystemNamespace = false
		logger.Warning("[AccessKubeSystemNamespace] DENIED")
	}else {
		profile.AccessKubeSystemNamespace = true
		logger.Success("[AccessKubeSystemNamespace] GRANTED")
	}
	return profile, nil
}

