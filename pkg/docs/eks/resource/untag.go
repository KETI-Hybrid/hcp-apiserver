package resource

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type UntagResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type ResourceUntagtInput struct {
	ResourceArn string   `json:"resourceArn"`
	TagKeys     []string `json:"tagKeys"`
}

func (UntagResource) Uri() string {
	return "/eks/resource/untag"
}
func (UntagResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(ResourceUntagtInput{}, eks.UntagResourceOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
