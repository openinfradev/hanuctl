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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/go-logr/logr"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/utils/openstack/clientconfig"
	"gopkg.in/yaml.v2"

	"sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/services/compute"
)

// clusterCmd represents the cluster command
var deleteClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cluster delete called")
		deleteCluster()
	},
}

func deleteCluster() error {
	var clouds clientconfig.Clouds
	var cloud clientconfig.Cloud
	cloudsyaml := "/tmp/tacoctl/clouds.yaml"
	content, err := ioutil.ReadFile(cloudsyaml)
	if err != nil {
	        return err
	}
	err = yaml.Unmarshal(content, &clouds)
	if err != nil {
        	return fmt.Errorf("failed to unmarshal clouds credentials stored in %v", cloudsyaml)
	}
	cloud = clouds.Clouds["taco-openstack"]
	cacert := "dummy"
	osProviderClient, clientOpts, err := newClient(cloud, []byte(cacert))
	if err != nil {
        	return err
	}
	var logger logr.Logger
	computeService, err := compute.NewService(osProviderClient, clientOpts, logger)
	computeClient, err := openstack.NewComputeV2(osProviderClient, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return fmt.Errorf("failed to create compute service client: %v", err)
	}
	opts := &compute.InstanceListOpts{
        	Name:  "taco-cluster",
	}
	instanceList, err := computeService.GetInstanceList(opts)
	if err != nil {
	        return err
	}
	if len(instanceList) == 0 {
        	return nil
	}
	fmt.Println("These instances will be deleted.")
	for _, instance := range instanceList {
		fmt.Println("instance id:",instance.ID)
		servers.Delete(computeClient, instance.ID).ExtractErr()
	}

	return nil
}

func newClient(cloud clientconfig.Cloud, caCert []byte) (*gophercloud.ProviderClient, *clientconfig.ClientOpts, error) {
        clientOpts := new(clientconfig.ClientOpts)
        if cloud.AuthInfo != nil {
                clientOpts.AuthInfo = cloud.AuthInfo
                clientOpts.AuthType = cloud.AuthType
                clientOpts.Cloud = cloud.Cloud
                clientOpts.RegionName = cloud.RegionName
        }

        opts, err := clientconfig.AuthOptions(clientOpts)
        if err != nil {
                return nil, nil, err
        }
        opts.AllowReauth = true

        provider, err := openstack.NewClient(opts.IdentityEndpoint)
        if err != nil {
                return nil, nil, fmt.Errorf("create providerClient err: %v", err)
        }

        config := &tls.Config{
                RootCAs: x509.NewCertPool(),
        }
        if cloud.Verify != nil {
                config.InsecureSkipVerify = !*cloud.Verify
        }
        config.RootCAs.AppendCertsFromPEM(caCert)

        provider.HTTPClient.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment, TLSClientConfig: config}
        err = openstack.Authenticate(provider, *opts)
        if err != nil {
                return nil, nil, fmt.Errorf("providerClient authentication err: %v", err)
        }
        return provider, clientOpts, nil
}


func init() {
	deleteCmd.AddCommand(deleteClusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
