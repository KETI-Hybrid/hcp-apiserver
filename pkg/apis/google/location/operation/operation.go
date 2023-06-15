package operation

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func LocationOperationsResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(CancelResource))
	apis.AddResource(router, new(GetResource))
	apis.AddResource(router, new(ListResource))
}
