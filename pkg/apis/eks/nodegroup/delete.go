package nodegroup

import (
	"hcp-apiserver/pkg/apis"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DeleteResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

func (DeleteResource) Uri() string {
	return "/eks/nodegroup/delete"
}
func (DeleteResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	return apis.Response{Code: 200, Data: nil}
}
