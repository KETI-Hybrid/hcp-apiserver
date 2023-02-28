package operation

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type CancelResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.GetNotSupported
}

type Cancel struct {
	ProjectName   string `json:"projectName"`
	LocationName  string `json:"locationName"`
	Operationname string `json:"operationName"`
}

func (CancelResource) Uri() string {
	return "/gke/zone/operation/cancel"
}

func (CancelResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Cancel{}, container.Empty{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
