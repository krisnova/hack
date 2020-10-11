package explorer

import (
	"github.com/kris-nova/logger"
	"os"
	"os/signal"
	"syscall"
)

var BreakNow = false

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