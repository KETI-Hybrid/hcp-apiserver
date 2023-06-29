package update

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type DescribeResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type UpdateDescribeInput struct {
	AddonName     string `json:"addonName"`
	Name          string `json:"name"`
	NodegroupName string `json:"nodegroupName"`
	UpdateId      string `json:"updateId"`
}

func (DescribeResource) Uri() string {
	return "/eks/update/describe"
}
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(UpdateDescribeInput{}, eks.DescribeUpdateOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
