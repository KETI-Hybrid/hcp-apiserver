package cluster

import (
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func ClusterResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CreateResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(ListResource))
}
