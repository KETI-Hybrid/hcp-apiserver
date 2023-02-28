package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type UpdateResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type Update struct {
	ProjectName  string                   `json:"projectName"`
	LocationName string                   `json:"locationName"`
	ClusterName  string                   `json:"clusterName"`
	Update       *container.ClusterUpdate `json:"update"`
}

func (UpdateResource) Uri() string {
	return "/gke/zone/cluster/update"
}

func (UpdateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Update{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
