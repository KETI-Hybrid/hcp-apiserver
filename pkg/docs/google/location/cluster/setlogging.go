package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type SetLoggingResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type SetLogging struct {
	ProjectName    string `json:"projectName"`
	LocationName   string `json:"locationName"`
	ClusterName    string `json:"clusterName"`
	LoggingService string `json:"loggingService"`
}

func (SetLoggingResource) Uri() string {
	return "/gke/locations/cluster/setLogging"
}

func (SetLoggingResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(SetLogging{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
