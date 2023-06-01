package managedcluster

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var ManagedClustersClient *armcontainerservice.ManagedClustersClient

func ManagedClusterResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateOrUpdateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(GetAccessProfileResource))
	apis.AddResource(router, new(GetCommandResultResource))
	apis.AddResource(router, new(GetOSOptionsResource))
	apis.AddResource(router, new(GetUpgradeProfileResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(ListByResourceGroupResource))
	apis.AddResource(router, new(ListClusterAdminCredentialsResource))
	apis.AddResource(router, new(ListClusterMonitoringUserCredentialsResource))
	apis.AddResource(router, new(ListClusterUserCredentialsResource))
	apis.AddResource(router, new(ListOutboundNetworkDependenciesEndpointsResource))
	apis.AddResource(router, new(ResetAADProfileResource))
	apis.AddResource(router, new(ResetServicePrincipalProfileResource))
	apis.AddResource(router, new(RotateClusterCertificatesResource))
	apis.AddResource(router, new(RotateServiceAccountSigningKeysResource))
	apis.AddResource(router, new(RunCommandResource))
	apis.AddResource(router, new(StartResource))
	apis.AddResource(router, new(StopResource))
	apis.AddResource(router, new(UpdateTagsResource))
	managedClusterClientInit()
}

func managedClusterClientInit() {
	sess := types.GetAKSClient()
	ManagedClustersClient = sess.ManagedClusterClient
}
