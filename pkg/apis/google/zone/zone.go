package zone

import (
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/apis/google/zone/cluster"
	"hcp-apiserver/pkg/apis/google/zone/operation"

	"github.com/julienschmidt/httprouter"
)

func ZoneResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(GetServerConfigResource))
	cluster.ZoneClusterNodepoolResourceAttach(router)
	operation.ZoneOperationsResourceAttach(router)
}
