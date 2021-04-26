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
	"github.com/kris-nova/hack/explorer"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

// inCmd represents the in command
var inCmd = &cobra.Command{
	Use:   "in",
	Short: "Tools for inside a cluster. Based on 'InCluster' client-go",
	Long: `

Now that we are inside a cluster let's see what we can do.'

`,
	Run: func(cmd *cobra.Command, args []string) {
		err := InCluster(inopt)
		if err != nil {
			logger.Critical(err.Error())
			os.Exit(2)
		}
		logger.Always("Exit...")
		os.Exit(0)
	},
}

var inopt = &explorer.LocalExplorerOptions{}

func init() {
	rootCmd.AddCommand(inCmd)
	inCmd.Flags().StringVarP(&inopt.Namespace, "namespace", "n", "default", "The namespace in Kubernetes to attach to.")
}

func InCluster(inopt *explorer.LocalExplorerOptions ) error {
	logger.Always("In Cluster...")
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	localeEx := explorer.NewLocalExplorer(*clientset, *config, inopt)
	err = localeEx.Probe()
	if err != nil {
		return err
	}
	return nil
}


