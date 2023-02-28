package addon

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type DescribeConfigResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type AddonDescribeConfigInput struct {
	AddonName    string `json:"addonName"`
	AddonVersion string `json:"addonVersion"`
}

func (DescribeConfigResource) Uri() string {
	return "/eks/addon/describe-confog"
}
func (DescribeConfigResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(AddonDescribeConfigInput{}, eks.DescribeAddonConfigurationOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
