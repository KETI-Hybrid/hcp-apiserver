package identityproviderconfig

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type DisassociateResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type IdentityProviderConfigDisassociateInput struct {
	ClientRequestToken     string                 `json:"clientRequestToken"`
	ClusterName            string                 `json:"clusterName"`
	IdentityProviderConfig IdentityProviderConfig `json:"identityProviderConfig"`
}

func (DisassociateResource) Uri() string {
	return "/eks/identity-provider-config/disassociate"
}
func (DisassociateResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(IdentityProviderConfigDisassociateInput{}, eks.DisassociateIdentityProviderConfigOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
