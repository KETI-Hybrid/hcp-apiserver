package wellknown

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func LocationClusterWellKnownResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(GetOpenidconfigurationResource))
}
