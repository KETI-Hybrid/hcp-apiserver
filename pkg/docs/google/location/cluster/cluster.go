package cluster

import (
	"hcp-apiserver/pkg/apis/google/location/cluster/nodepool"
	"hcp-apiserver/pkg/apis/google/location/cluster/wellknown"
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func LocationClusterResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CompleteIpRotationResource))
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(GetJwksResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(SetAddonsResource))
	docs.AddResource(router, new(SetLegacyAbacResource))
	docs.AddResource(router, new(SetLoggingResource))
	docs.AddResource(router, new(SetMaintenancePolicyResource))
	docs.AddResource(router, new(SetMasterAuthResource))
	docs.AddResource(router, new(SetMonitoringResource))
	docs.AddResource(router, new(SetNetworkPolicyResource))
	docs.AddResource(router, new(SetResourceLabelsResource))
	docs.AddResource(router, new(StartIpRotationResource))
	docs.AddResource(router, new(UpdateResource))
	docs.AddResource(router, new(UpdateMasterResource))
	nodepool.LocationClusterNodepoolResourceAttach(router)
	wellknown.LocationClusterWellKnownResourceAttach(router)
}
