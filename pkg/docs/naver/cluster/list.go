package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	"github.com/julienschmidt/httprouter"
)

type ListResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

func (ListResource) Uri() string {
	return "/nks/cluster/list"
}

func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	response := util.DocWithoutReq(vnks.ClustersRes{})

	resp := docs.ForDoc{
		Req:  nil,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
