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
	"fmt"

	"k8s.io/klog"
	"github.com/spf13/cobra"
	cluster "sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/clusterdeployer/bootstrap"
)

type CreateOptions struct {
        Cluster                 string
        Machine                 string
        ProviderComponents      string
        AddonComponents         string
        BootstrapOnlyComponents string
        Provider                string
        KubeconfigOutput        string
        BootstrapFlags          bootstrap.Options
}

var co = &cluster.CreateOptions{}

// clusterCmd represents the cluster command
var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cluster called")
		CreateCluster(co)
	},
}

func CreateCluster(co *cluster.CreateOptions)  {
	co.Cluster = "/tmp/tacoctl/cluster.yaml"
	co.Machine = "/tmp/tacoctl/controlplane.yaml"
	co.ProviderComponents = "/tmp/tacoctl/provider-components.yaml"
	co.AddonComponents = "/tmp/tacoctl/addons.yaml"
	co.BootstrapFlags.KubeConfig = "/tmp/tacoctl/kubeconfig"
	co.BootstrapFlags.Type = "none"
	co.KubeconfigOutput = "/tmp/tacoctl/targetconfig"
	if err := cluster.RunCreate(co); err != nil {
        	klog.Exit(err)
	}
}

func init() {
	// Required flags
	//createClusterCmd.Flags().StringVarP(&co.Cluster, "clusterconfig", "c", "", "A yaml file containing cluster object definition. Required.")
	//createClusterCmd.MarkFlagRequired("clusterconfig")
        //viper.BindPFlag("clusterconfig", createClusterCmd.Flags().Lookup("clusterconfig"))
	createCmd.AddCommand(createClusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
