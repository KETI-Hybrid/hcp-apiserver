package nodepools

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func NodePoolResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(DeleteResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(CreateResource))
	apis.AddResource(router, new(UpdateResource))
}
