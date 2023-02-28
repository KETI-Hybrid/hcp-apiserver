package wellknown

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type GetOpenidconfigurationResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

type GetOpenidconfiguration struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
}

func (GetOpenidconfigurationResource) Uri() string {
	return "/gke/locations/cluster/well-known/openid-configuration"
}

func (GetOpenidconfigurationResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(GetOpenidconfiguration{}, container.GetOpenIDConfigResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
