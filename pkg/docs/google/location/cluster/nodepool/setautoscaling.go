package nodepool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type SetAutoscalingResource struct {
	docs.GetNotSupported
	docs.DeleteNotSupported
	docs.PutNotSupported
}

type Autoscaling struct {
	ProjectName  string                         `json:"projectName"`
	LocationName string                         `json:"locationName"`
	ClusterName  string                         `json:"clusterName"`
	NodePoolName string                         `json:"nodepoolName"`
	Autoscaling  *container.NodePoolAutoscaling `json:"autoscaling"`
}

func (SetAutoscalingResource) Uri() string {
	return "/gke/locations/cluster/nodePools/autoscaling"
}

func (SetAutoscalingResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Autoscaling{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
