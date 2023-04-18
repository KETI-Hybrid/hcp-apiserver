package k8sapi

import (
	"hcp-apiserver/pkg/apis"

	"github.com/julienschmidt/httprouter"
)

func KubernetesAPIResourceAttach(router *httprouter.Router) {
	apis.AddResource(router, new(AccessResource))
}
