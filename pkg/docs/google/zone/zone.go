package zone

import (
	"hcp-apiserver/pkg/apis/google/zone/cluster"
	"hcp-apiserver/pkg/apis/google/zone/operation"
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func ZoneResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(GetServerConfigResource))
	cluster.ZoneClusterNodepoolResourceAttach(router)
	operation.ZoneOperationsResourceAttach(router)
}
