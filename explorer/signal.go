package explorer

import (
	"context"
	"github.com/kris-nova/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"os/signal"
	"syscall"
)

var BreakNow = false
var Recieved = false

func HandleSignals() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for {
			s := <-sigCh
			switch s {
			case syscall.SIGHUP:
				logger.Warning("SIGHUP")
				logger.Warning("Shutting down...")
				BreakNow = true
			case syscall.SIGINT:
				logger.Critical("SIGINT")
				logger.Critical("Shutting down...")
				BreakNow = true
			case syscall.SIGTERM:
				logger.Critical("SIGTERM")
				logger.Critical("Shutting down...")
				BreakNow = true
			case syscall.SIGQUIT:
				logger.Warning("SIGHUP")
				logger.Warning("Shutting down...")
				BreakNow = true
			default:
				logger.Info("UNKNOWN SIGNAL")
			}
		}
	}()
}

func (e *RemoteExplorer) DeferDeletePod(name, namespace string) {
	go func() {
		for !BreakNow {}
		logger.Always("Deleting pod: %s", name)
		err := e.ClientSet.CoreV1().Pods(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
		if err != nil {
			logger.Warning("Error deleting pod: %v", err)
		}
	}()

}