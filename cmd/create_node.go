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
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/clusterdeployer/clusterclient"
	//"sigs.k8s.io/cluster-api/util/yaml"
)

type ApplyMachineDeploymentsOptions struct {
        Kubeconfig string
        MachineDeployments   string
        Namespace  string
}

var amdo = &ApplyMachineDeploymentsOptions{}


// createNodeCmd represents the node command
var createNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApplyMachines(amdo)
	},
}

func ApplyMachines(amdo *ApplyMachineDeploymentsOptions) error{
	amdo.Kubeconfig = "/tmp/hanuctl/targetconfig"
	amdo.MachineDeployments = "/tmp/hanuctl/machinedeployment.yaml"
	amdo.Namespace = "taco-cluster"

        kubeconfig, err := ioutil.ReadFile(amdo.Kubeconfig)
        if err != nil {
                return err
        }

//        out, err := yaml.Parse(yaml.ParseInput{File: amdo.MachineDeployments})
//        if err != nil {
//                return err
//        }
//
	out, err := ioutil.ReadFile(amdo.MachineDeployments)
	if err != nil {
		return err
	}
        clientFactory := clusterclient.NewFactory()
        client, err := clientFactory.NewClientFromKubeconfig(string(kubeconfig))
	if err := client.Apply(string(out)); err != nil {
		return errors.Wrap(err, "unable to apply machine deployments")
	}
//        if err != nil {
//                return errors.Wrap(err, "unable to create cluster client")
//        }
//
//        err = client.EnsureNamespace(amdo.Namespace)
//        if err != nil {
//                return errors.Wrapf(err, "unable to ensure namespace %q", amdo.Namespace)
//        }
//
//        if err := client.CreateMachineDeployments(out.MachineDeployments, amdo.Namespace); err != nil {
//                return errors.Wrap(err, "unable to apply machinedeployments")
//        }
	
        return nil
}

func init() {
	createCmd.AddCommand(createNodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createNodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createNodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
