package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type ListResource struct {
	docs.PostNotSupported
	docs.DeleteNotSupported
	docs.PutNotSupported
}

type List struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
}

func (ListResource) Uri() string {
	return "/gke/zone/cluster/list"
}

func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(List{}, container.ListClustersResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
