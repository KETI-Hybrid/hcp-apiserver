package snapshot

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type ListResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.PutNotSupported
}

func (ListResource) Uri() string {
	return "/aks/snapshot/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	response := util.DocWithoutReq(armcontainerservice.SnapshotsClientListResponse{})

	resp := docs.ForDoc{
		Req:  nil,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
