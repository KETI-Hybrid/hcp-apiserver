package cluster

import (
	"hcp-apiserver/pkg/apis"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RegisterResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

func (RegisterResource) Uri() string {
	return "/eks/cluster/register"
}
func (RegisterResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	return apis.Response{Code: 200, Data: nil}
}
