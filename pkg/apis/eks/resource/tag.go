package resource

import (
	"hcp-apiserver/pkg/apis"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TagResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

func (TagResource) Uri() string {
	return "/eks/resource/tag"
}
func (TagResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	return apis.Response{Code: 200, Data: nil}
}
