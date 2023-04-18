package cluster

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func ClusterResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(DeregisterResource))
	apis.AddResource(router, new(RegisterResource))
	apis.AddResource(router, new(DescribeResource))
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(UpgradeResource))
}
