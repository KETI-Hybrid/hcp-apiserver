package k8sapi

import (
	"hcp-apiserver/pkg/docs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AccessResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

func (AccessResource) Uri() string {
	return "/eks/kubernetes-api/access"
}
func (AccessResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	return docs.Response{Code: 200, Data: nil}
}
