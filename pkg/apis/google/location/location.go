package location

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/apis/google/location/cluster"
	"hcp-apiserver/pkg/apis/google/location/operation"

	"github.com/julienschmidt/httprouter"
)

func LocationResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(GetServerConfigResource))
	cluster.LocationClusterResourceAttach(router)
	operation.LocationOperationsResourceAttach(router)
}
