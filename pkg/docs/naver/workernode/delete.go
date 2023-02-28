package workernode

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
	NodePoolID  string `json:"nodePoolID"`
	NodeNumber  string `json:"nodeNumber"`
}

type DeleteResp struct {
	ClusterUUID string `json:"clusteruuid"`
	NodePoolID  string `json:"nodePoolID"`
	NodeNumber  string `json:"nodeNumber"`
	Status      bool   `json:"status"`
}

func (DeleteResource) Uri() string {
	return "/nks/workerNode/delete"
}

func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Delete{}, DeleteResp{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
