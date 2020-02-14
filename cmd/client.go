package cmd

import (
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/services/compute"
)

type client struct {
	Service		*compute.Service
}


