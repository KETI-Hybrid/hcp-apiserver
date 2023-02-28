package addon

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

type AddonDescribeInput struct {
	AddonName   string `json:"addonName"`
	ClusterName string `json:"clusterName"`
}

func (DescribeResource) Uri() string {
	return "/eks/addon/describe"
}
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(AddonDescribeInput{}, eks.DescribeAddonOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
