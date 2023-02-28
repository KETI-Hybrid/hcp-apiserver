package workernode

import (
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func WorkerNodeResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
}
