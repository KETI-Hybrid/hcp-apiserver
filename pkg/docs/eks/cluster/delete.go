package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type DeleteResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.GetNotSupported
}

type ClusterDeleteInput struct {
	Name string `json:"name"`
}

func (DeleteResource) Uri() string {
	return "/eks/cluster/delete"
}
func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(ClusterDeleteInput{}, eks.DeleteClusterOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
