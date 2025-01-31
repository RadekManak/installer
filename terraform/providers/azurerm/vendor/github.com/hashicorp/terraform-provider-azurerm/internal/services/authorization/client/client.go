package client

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2020-04-01-preview/authorization" // nolint: staticcheck // nolint: staticcheck
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentscheduleinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentschedulerequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityscheduleinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityschedulerequests"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	RoleAssignmentsClient                  *authorization.RoleAssignmentsClient
	RoleDefinitionsClient                  *authorization.RoleDefinitionsClient
	RoleAssignmentScheduleRequestClient    *roleassignmentschedulerequests.RoleAssignmentScheduleRequestsClient
	RoleAssignmentScheduleInstancesClient  *roleassignmentscheduleinstances.RoleAssignmentScheduleInstancesClient
	RoleEligibilityScheduleRequestClient   *roleeligibilityschedulerequests.RoleEligibilityScheduleRequestsClient
	RoleEligibilityScheduleInstancesClient *roleeligibilityscheduleinstances.RoleEligibilityScheduleInstancesClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	roleAssignmentsClient := authorization.NewRoleAssignmentsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&roleAssignmentsClient.Client, o.ResourceManagerAuthorizer)

	roleDefinitionsClient := authorization.NewRoleDefinitionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&roleDefinitionsClient.Client, o.ResourceManagerAuthorizer)

	roleAssignmentScheduleRequestsClient, err := roleassignmentschedulerequests.NewRoleAssignmentScheduleRequestsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("creating roleAssignmentScheduleRequestsClient: %+v", err)
	}

	o.Configure(roleAssignmentScheduleRequestsClient.Client, o.Authorizers.ResourceManager)

	roleAssignmentScheduleInstancesClient, err := roleassignmentscheduleinstances.NewRoleAssignmentScheduleInstancesClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("creating roleAssignmentScheduleInstancesClient: %+v", err)
	}
	o.Configure(roleAssignmentScheduleInstancesClient.Client, o.Authorizers.ResourceManager)

	roleEligibilityScheduleRequestClient, err := roleeligibilityschedulerequests.NewRoleEligibilityScheduleRequestsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("creating roleEligibilityScheduleRequestClient: %+v", err)
	}
	o.Configure(roleEligibilityScheduleRequestClient.Client, o.Authorizers.ResourceManager)

	roleEligibilityScheduleInstancesClient, err := roleeligibilityscheduleinstances.NewRoleEligibilityScheduleInstancesClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("creating roleEligibilityScheduleInstancesClient: %+v", err)
	}
	o.Configure(roleEligibilityScheduleInstancesClient.Client, o.Authorizers.ResourceManager)

	return &Client{
		RoleAssignmentsClient:                  &roleAssignmentsClient,
		RoleDefinitionsClient:                  &roleDefinitionsClient,
		RoleAssignmentScheduleRequestClient:    roleAssignmentScheduleRequestsClient,
		RoleAssignmentScheduleInstancesClient:  roleAssignmentScheduleInstancesClient,
		RoleEligibilityScheduleRequestClient:   roleEligibilityScheduleRequestClient,
		RoleEligibilityScheduleInstancesClient: roleEligibilityScheduleInstancesClient,
	}, nil
}
