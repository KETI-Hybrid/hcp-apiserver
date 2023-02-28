package nodepool

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func LocationClusterNodepoolResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CompleteUpgradeResource))
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(RollbackResource))
	apis.AddResource(router, new(SetAutoscalingResource))
	apis.AddResource(router, new(SetManagementResource))
	apis.AddResource(router, new(SetSizeResource))
	apis.AddResource(router, new(UpdateResource))
}
