package azure

import (
	"hcp-apiserver/pkg/apis/azure/agentpool"
	"hcp-apiserver/pkg/apis/azure/maintenanceconfig"
	"hcp-apiserver/pkg/apis/azure/managedcluster"
	"hcp-apiserver/pkg/apis/azure/operation"
	"hcp-apiserver/pkg/apis/azure/privateendpoint"
	"hcp-apiserver/pkg/apis/azure/privatelink"
	"hcp-apiserver/pkg/apis/azure/resolveprivatelink"
	"hcp-apiserver/pkg/apis/azure/snapshot"

	"github.com/julienschmidt/httprouter"
)

func InitAKSEndPoint(router *httprouter.Router) {
	agentpool.AgentPoolResourceAttach(router)
	maintenanceconfig.MaintenanceConfigurationResourceAttach(router)
	managedcluster.ManagedClusterResourceAttach(router)
	operation.OperationsResourceAttach(router)
	privateendpoint.PrivateEndpointConnectionsResourceAttach(router)
	privatelink.PrivateLinkResourceAttach(router)
	resolveprivatelink.ResolvePrivateLinkServiceIdResourceAttach(router)
	snapshot.SnapshotResourceAttach(router)
}
