package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type AddonResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type Addon struct {
	ProjectName  string                  `json:"projectName"`
	LocationName string                  `json:"locationName"`
	ClusterName  string                  `json:"clusterName"`
	AddonConfig  *container.AddonsConfig `json:"addonsConfig"`
}

func (AddonResource) Uri() string {
	return "/gke/zone/cluster/addon"
}

func (AddonResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Addon{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
