package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type SetMasterAuthResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type SetMasterAuth struct {
	ProjectName  string                `json:"projectName"`
	LocationName string                `json:"locationName"`
	ClusterName  string                `json:"clusterName"`
	Action       string                `json:"action"`
	Update       *container.MasterAuth `json:"update"`
}

func (SetMasterAuthResource) Uri() string {
	return "/gke/zone/cluster/setMasterAuth"
}

func (SetMasterAuthResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(SetMasterAuth{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
