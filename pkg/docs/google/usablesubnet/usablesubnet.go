package usablesubnet

import (
	"hcp-apiserver/pkg/docs"

	"github.com/julienschmidt/httprouter"
)

func UsableSubnetworksResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(ListResource))
}
