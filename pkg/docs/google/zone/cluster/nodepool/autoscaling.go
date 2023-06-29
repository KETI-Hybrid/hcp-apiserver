package nodepool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type AutoScalingResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type AutoScaling struct {
	ProjectName  string                         `json:"projectName"`
	LocationName string                         `json:"locationName"`
	ClusterName  string                         `json:"clusterName"`
	NodePoolName string                         `json:"nodePoolName"`
	NodePool     *container.NodePoolAutoscaling `json:"nodePool"`
}

func (AutoScalingResource) Uri() string {
	return "/gke/zone/cluster/nodePools/autoscaling"
}

func (AutoScalingResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(AutoScaling{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
