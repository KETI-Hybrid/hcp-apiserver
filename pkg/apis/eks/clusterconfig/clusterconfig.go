package clusterconfig

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func ClusterConfigResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(UpdateResource))
}
