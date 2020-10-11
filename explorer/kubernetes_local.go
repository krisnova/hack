package explorer

import (
	"context"
	"github.com/kris-nova/logger"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type LocalExplorer struct {
	ClientSet kubernetes.Clientset
	Config rest.Config
	Options *LocalExplorerOptions
}

type LocalExplorerOptions struct {
	Namespace string
}

func NewLocalExplorer( clientSet kubernetes.Clientset, config rest.Config, options *LocalExplorerOptions) *LocalExplorer{
	return &LocalExplorer{
		ClientSet: clientSet,
		Config: config,
		Options: options,
	}
}

func (l *LocalExplorer) Probe() error {
	/**
	 * TODO @kris-nova
	 * Now that we are here we should have a pre-loaded container with all our tools
	 * And we should have as many privileges as possible
	 * We can try container escaping and exploring namespaces in the system
	 * We also have linux capabilities to explore and we can see what we can do in /proc
	 *
	 * During the probe we also need to set up as many RBAC permissions as possible and
	 * show all relevant kubectl commands for them
	 */

	 /**
	  * For now let's just list pods
	  */
	pods, err := l.ClientSet.CoreV1().Pods(l.Options.Namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return err
	}
	for _, pod := range pods.Items {
		logger.Always("Pod: %s", pod.Name)
	}
	return nil
}

/**
 *
 * TODO @kris-nova Can we build a set of unit-test like checks for seeing what is and isn't vulnerable in the cluster?
 *
 */