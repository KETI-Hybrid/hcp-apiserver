package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type UpdateMasterResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type UpdateMaster struct {
	ProjectName   string `json:"projectName"`
	LocationName  string `json:"locationName"`
	ClusterName   string `json:"clusterName"`
	MasterVersion string `json:"masterVersion"`
}

func (UpdateMasterResource) Uri() string {
	return "/gke/locations/cluster/updateMaster"
}

func (UpdateMasterResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(UpdateMaster{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
