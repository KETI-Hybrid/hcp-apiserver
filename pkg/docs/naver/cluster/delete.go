package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DeleteResource struct {
	docs.PutNotSupported
	docs.PostNotSupported
	docs.GetNotSupported
}

type Delete struct {
	ClusterUUID string `json:"clusteruuid"`
}

type DeleteResp struct {
	ClusterUUID string `json:"clusteruuid"`
	Status      bool   `json:"status"`
}

func (DeleteResource) Uri() string {
	return "/nks/cluster/delete"
}

func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Delete{}, DeleteResp{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
