package addon

import (
	"hcp-apiserver/pkg/apis"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ListResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

func (ListResource) Uri() string {
	return "/eks/addon/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	return apis.Response{Code: 200, Data: nil}
}
