package nodegroupconfig

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func NodeGroupConfigResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(UpdateResource))
}
