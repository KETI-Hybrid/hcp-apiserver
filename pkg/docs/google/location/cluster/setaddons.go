package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type SetAddonsResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type SetAddons struct {
	ProjectName  string                  `json:"projectName"`
	LocationName string                  `json:"locationName"`
	ClusterName  string                  `json:"clusterName"`
	AddonsConfig *container.AddonsConfig `json:"addonConfig"`
}

func (SetAddonsResource) Uri() string {
	return "/gke/locations/cluster/setAddon"
}

func (SetAddonsResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(SetAddons{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
