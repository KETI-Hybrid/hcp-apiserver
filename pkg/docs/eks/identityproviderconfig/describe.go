package identityproviderconfig

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

type IdentityProviderConfigDescribeInput struct {
	ClusterName            string                 `json:"clusterName"`
	IdentityProviderConfig IdentityProviderConfig `json:"identityProviderConfig"`
}

type IdentityProviderConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (DescribeResource) Uri() string {
	return "/eks/identity-provider-config/describe"
}
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(IdentityProviderConfigDescribeInput{}, eks.DescribeIdentityProviderConfigOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
