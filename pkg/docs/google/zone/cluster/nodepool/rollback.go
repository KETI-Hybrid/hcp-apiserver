package nodepool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type RollbackResource struct {
	docs.GetNotSupported
	docs.DeleteNotSupported
	docs.PutNotSupported
}

type Rollback struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
	NodePoolName string `json:"nodepoolName"`
	RespectPdb   bool   `json:"respectPdb"`
}

func (RollbackResource) Uri() string {
	return "/gke/zone/cluster/nodePools/rollback"
}

func (RollbackResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Rollback{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
