package operation

import (
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func ZoneOperationsResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(CancelResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(ListResource))
}
