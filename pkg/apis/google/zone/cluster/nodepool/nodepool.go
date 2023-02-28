package nodepool

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func ZoneClusterNodepoolResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(AutoScalingResource))
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(RollbackResource))
	apis.AddResource(router, new(SetManagementResource))
	apis.AddResource(router, new(SetSizeResource))
	apis.AddResource(router, new(UpdateResource))
}
