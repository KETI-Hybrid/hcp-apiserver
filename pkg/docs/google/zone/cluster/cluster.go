package cluster

import (
	"hcp-apiserver/pkg/apis/google/zone/cluster/nodepool"
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func ZoneClusterNodepoolResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(AddonResource))
	docs.AddResource(router, new(CompleteIpRotationResource))
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(LegacyAbacResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(LoggingResource))
	docs.AddResource(router, new(MasterResource))
	docs.AddResource(router, new(MonitoringResource))
	docs.AddResource(router, new(SetMaintenancePolicyResource))
	docs.AddResource(router, new(SetMasterAuthResource))
	docs.AddResource(router, new(SetNetworkPolicyResource))
	docs.AddResource(router, new(StartIpRotationResource))
	docs.AddResource(router, new(UpdateResource))

	nodepool.ZoneClusterNodepoolResourceAttach(router)
}
