package client

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-11-10/machineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-11-10/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-11-10/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-11-10/privatelinkscopes"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	MachineExtensionsClient          *machineextensions.MachineExtensionsClient
	MachinesClient                   *machines.MachinesClient
	PrivateEndpointConnectionsClient *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkScopesClient          *privatelinkscopes.PrivateLinkScopesClient
}

func NewClient(o *common.ClientOptions) *Client {

	machineExtensionsClient := machineextensions.NewMachineExtensionsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&machineExtensionsClient.Client, o.ResourceManagerAuthorizer)

	machinesClient := machines.NewMachinesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&machinesClient.Client, o.ResourceManagerAuthorizer)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&privateEndpointConnectionsClient.Client, o.ResourceManagerAuthorizer)

	privateLinkScopesClient := privatelinkscopes.NewPrivateLinkScopesClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&privateLinkScopesClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		MachineExtensionsClient:          &machineExtensionsClient,
		MachinesClient:                   &machinesClient,
		PrivateEndpointConnectionsClient: &privateEndpointConnectionsClient,
		PrivateLinkScopesClient:          &privateLinkScopesClient,
	}
}
