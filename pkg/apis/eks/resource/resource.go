package resource

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func ResourceResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(ListResource))
	apis.AddResource(router, new(TagResource))
	apis.AddResource(router, new(UntagResource))
}
