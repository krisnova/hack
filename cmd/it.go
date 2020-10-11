/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/kris-nova/hack/explorer"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
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
		opt.KubeconfigPath = workingPath // Update the manipulated path
		err := InteractiveTTY(opt)
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
	itCmd.Flags().IntVarP(&logger.Level, "verbosity", "v", 4, "Verbosity level")
	itCmd.Flags().StringVarP(&opt.KubeconfigPath, "kubeconfig", "k", "~/.kube/config", "The path on the local filesystem to read a Kube config file from.")
}

type ITOptions struct {
	KubeconfigPath string
}

var opt = &ITOptions{}

/**
 * InteractiveTTY is an opinionated way to enter a pod in a Kubernetes cluster.
 */
func InteractiveTTY(opt *ITOptions) error {
	config, err := clientcmd.BuildConfigFromFlags("", opt.KubeconfigPath)
	if err != nil {
		return err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	remoteEx := explorer.NewRemoteExplorer(*clientSet)
	_, err = remoteEx.Probe()
	if err != nil {
		return err
	}


	return nil
}