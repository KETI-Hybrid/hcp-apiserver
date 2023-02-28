package nodepools

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	"github.com/julienschmidt/httprouter"
)

type GetResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

type Get struct {
	ClusterUUID string `json:"clusteruuid"`
}

func (GetResource) Uri() string {
	return "/nks/nodepool/get"
}

func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Get{}, vnks.NodePoolsRes{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
