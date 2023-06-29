package nodepools

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	"github.com/julienschmidt/httprouter"
)

type CreateResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.GetNotSupported
}

type Create struct {
	ClusterUUID string                     `json:"clusteruuid"`
	NodePool    *vnks.NodePoolCreationBody `json:"nodepool"`
}

type CreateResp struct {
	ClusterUUID  string `json:"clusteruuid"`
	NodePoolName string `json:"nodepoolName"`
	Status       bool   `json:"status"`
}

func (CreateResource) Uri() string {
	return "/nks/nodepool/create"
}

func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Create{}, CreateResp{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
