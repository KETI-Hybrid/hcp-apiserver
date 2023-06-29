package nodepool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type CreateResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type Create struct {
	ProjectName  string              `json:"projectName"`
	LocationName string              `json:"locationName"`
	ClusterName  string              `json:"clusterName"`
	NodePool     *container.NodePool `json:"nodepool"`
}

func (CreateResource) Uri() string {
	return "/gke/locations/cluster/nodePools/create"
}

func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Create{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
