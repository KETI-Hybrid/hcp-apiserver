package wellknown

import (
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func LocationClusterWellKnownResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(GetOpenidconfigurationResource))
}
