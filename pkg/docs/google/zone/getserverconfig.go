package zone

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type GetServerConfigResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

type GetServerConfig struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
}

func (GetServerConfigResource) Uri() string {
	return "/gke/zone/getServerconfig"
}

func (GetServerConfigResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(GetServerConfig{}, container.ServerConfig{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
