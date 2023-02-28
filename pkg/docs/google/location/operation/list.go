package operation

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type ListResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

type List struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
}

func (ListResource) Uri() string {
	return "/gke/locations/operation/list"
}

func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Cancel{}, container.ListOperationsResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
