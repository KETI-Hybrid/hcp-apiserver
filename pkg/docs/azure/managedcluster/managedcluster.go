package managedcluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var ManagedClustersClient *armcontainerservice.ManagedClustersClient

func ManagedClusterResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateOrUpdateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(GetAccessProfileResource))
	docs.AddResource(router, new(GetCommandResultResource))
	docs.AddResource(router, new(GetOSOptionsResource))
	docs.AddResource(router, new(GetUpgradeProfileResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(ListByResourceGroupResource))
	docs.AddResource(router, new(ListClusterAdminCredentialsResource))
	docs.AddResource(router, new(ListClusterMonitoringUserCredentialsResource))
	docs.AddResource(router, new(ListClusterUserCredentialsResource))
	docs.AddResource(router, new(ListOutboundNetworkDependenciesEndpointsResource))
	docs.AddResource(router, new(ResetAADProfileResource))
	docs.AddResource(router, new(ResetServicePrincipalProfileResource))
	docs.AddResource(router, new(RotateClusterCertificatesResource))
	docs.AddResource(router, new(RotateServiceAccountSigningKeysResource))
	docs.AddResource(router, new(RunCommandResource))
	docs.AddResource(router, new(StartResource))
	docs.AddResource(router, new(StopResource))
	docs.AddResource(router, new(UpdateTagsResource))
	managedClusterClientInit()
}

func managedClusterClientInit() {
	sess := types.GetAKSClient()
	ManagedClustersClient = sess.ManagedClusterClient
}
