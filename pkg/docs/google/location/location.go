package location

import (
	"hcp-apiserver/pkg/apis/google/location/cluster"
	"hcp-apiserver/pkg/apis/google/location/operation"
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func LocationResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(GetServerConfigResource))
	cluster.LocationClusterResourceAttach(router)
	operation.LocationOperationsResourceAttach(router)
}
