package usablesubnet

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func UsableSubnetworksResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(ListResource))
}
