package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type LoggingResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type Logging struct {
	ProjectName    string `json:"projectName"`
	LocationName   string `json:"locationName"`
	ClusterName    string `json:"clusterName"`
	LoggingService string `json:"loggingService"`
}

func (LoggingResource) Uri() string {
	return "/gke/zone/cluster/setLogging"
}

func (LoggingResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Logging{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
