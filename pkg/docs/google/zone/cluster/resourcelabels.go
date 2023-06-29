package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type ResourceLabelsResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type ResourceLabels struct {
	ProjectName      string            `json:"projectName"`
	LocationName     string            `json:"locationName"`
	ClusterName      string            `json:"clusterName"`
	ResourceLabels   map[string]string `json:"resourceLabels"`
	LabelFingerprint string            `json:"labelFingerprint"`
}

func (ResourceLabelsResource) Uri() string {
	return "/gke/zone/cluster/setResourceLabels"
}

func (ResourceLabelsResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(ResourceLabels{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
