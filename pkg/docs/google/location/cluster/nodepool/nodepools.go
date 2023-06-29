package nodepool

import (
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func LocationClusterNodepoolResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CompleteUpgradeResource))
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(RollbackResource))
	docs.AddResource(router, new(SetAutoscalingResource))
	docs.AddResource(router, new(SetManagementResource))
	docs.AddResource(router, new(SetSizeResource))
	docs.AddResource(router, new(UpdateResource))
}
