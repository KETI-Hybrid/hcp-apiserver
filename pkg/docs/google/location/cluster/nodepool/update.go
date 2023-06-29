package nodepool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type UpdateResource struct {
	docs.GetNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

type Update struct {
	ProjectName     string                     `json:"projectName"`
	LocationName    string                     `json:"locationName"`
	ClusterName     string                     `json:"clusterName"`
	NodePoolName    string                     `json:"nodepoolName"`
	NodeVersion     string                     `json:"nodeVersion,omitempty"`
	Locations       []string                   `json:"locations,omitempty"`
	ImageType       string                     `json:"imageType,omitempty"`
	UpgradeSettings *container.UpgradeSettings `json:"upgradeSettings,omitempty"`
}

func (UpdateResource) Uri() string {
	return "/gke/locations/cluster/nodePools/update"
}

func (UpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Update{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
