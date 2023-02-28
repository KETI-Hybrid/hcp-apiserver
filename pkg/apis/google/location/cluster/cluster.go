package cluster

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/apis/google/location/cluster/nodepool"
	"hcp-apiserver/pkg/apis/google/location/cluster/wellknown"

	"github.com/julienschmidt/httprouter"
)

func LocationClusterResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CompleteIpRotationResource))
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(GetJwksResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(SetAddonsResource))
	apis.AddResource(router, new(SetLegacyAbacResource))
	apis.AddResource(router, new(SetLoggingResource))
	apis.AddResource(router, new(SetMaintenancePolicyResource))
	apis.AddResource(router, new(SetMasterAuthResource))
	apis.AddResource(router, new(SetMonitoringResource))
	apis.AddResource(router, new(SetNetworkPolicyResource))
	apis.AddResource(router, new(SetResourceLabelsResource))
	apis.AddResource(router, new(StartIpRotationResource))
	apis.AddResource(router, new(UpdateResource))
	apis.AddResource(router, new(UpdateMasterResource))
	nodepool.LocationClusterNodepoolResourceAttach(router)
	wellknown.LocationClusterWellKnownResourceAttach(router)
}
