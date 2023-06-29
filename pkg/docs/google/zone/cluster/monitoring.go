package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type MonitoringResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type Monitoring struct {
	ProjectName       string `json:"projectName"`
	LocationName      string `json:"locationName"`
	ClusterName       string `json:"clusterName"`
	MonitoringService string `json:"monitoringService"`
}

func (MonitoringResource) Uri() string {
	return "/gke/zone/cluster/setMonitoring"
}

func (MonitoringResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Monitoring{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
