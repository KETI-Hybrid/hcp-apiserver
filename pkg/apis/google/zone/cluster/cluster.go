package cluster

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/apis/google/zone/cluster/nodepool"

	"github.com/julienschmidt/httprouter"
)

func ZoneClusterNodepoolResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(AddonResource))
	apis.AddResource(router, new(CompleteIpRotationResource))
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(LegacyAbacResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(LoggingResource))
	apis.AddResource(router, new(MasterResource))
	apis.AddResource(router, new(MonitoringResource))
	apis.AddResource(router, new(SetMaintenancePolicyResource))
	apis.AddResource(router, new(SetMasterAuthResource))
	apis.AddResource(router, new(SetNetworkPolicyResource))
	apis.AddResource(router, new(StartIpRotationResource))
	apis.AddResource(router, new(UpdateResource))

	nodepool.ZoneClusterNodepoolResourceAttach(router)
}
