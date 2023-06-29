package nodepools

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	"github.com/julienschmidt/httprouter"
)

type UpdateResource struct {
	docs.GetNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

type Update struct {
	ClusterUUID string                   `json:"clusteruuid"`
	NodeNumber  string                   `json:"nodeNumber"`
	UpdateBody  *vnks.NodePoolUpdateBody `json:"updateBody"`
}

type UpdateResp struct {
	ClusterUUID string `json:"clusteruuid"`
	NodeNumber  string `json:"nodeNumber"`
	Status      bool   `json:"status"`
}

func (UpdateResource) Uri() string {
	return "/nks/nodepool/update"
}

func (UpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Update{}, UpdateResp{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
