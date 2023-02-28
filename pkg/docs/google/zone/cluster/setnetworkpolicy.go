package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type SetNetworkPolicyResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type SetNetworkPolicy struct {
	ProjectName   string                   `json:"projectName"`
	LocationName  string                   `json:"locationName"`
	ClusterName   string                   `json:"clusterName"`
	NetworkPolicy *container.NetworkPolicy `json:"networkPolicy"`
}

func (SetNetworkPolicyResource) Uri() string {
	return "/gke/zone/cluster/setNetworkPolicy"
}

func (SetNetworkPolicyResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(SetNetworkPolicy{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
