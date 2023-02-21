package addon

import (
	"hcp-apiserver/pkg/apis"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DescribeResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

func (DescribeResource) Uri() string {
	return "/eks/addon/describe"
}
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	return apis.Response{Code: 200, Data: nil}
}
