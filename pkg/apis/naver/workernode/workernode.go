package workernode

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func WorkerNodeResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
}
