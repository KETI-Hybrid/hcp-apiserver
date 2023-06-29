package identityproviderconfig

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type AssociateResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type IdentityProviderConfigAssociateInput struct {
	ClientRequestToken string                            `json:"clientRequestToken"`
	ClusterName        string                            `json:"clusterName"`
	Oidc               OidcIdentityProviderConfigRequest `json:"oidc"`
	Tags               map[string]string                 `json:"tags"`
}

type OidcIdentityProviderConfigRequest struct {
	ClientId                   string            `json:"clientId"`
	GroupsClaim                string            `json:"groupsClaim"`
	GroupsPrefix               string            `json:"groupsPrefix"`
	IdentityProviderConfigName string            `json:"identityProviderConfigName"`
	IssuerUrl                  string            `json:"issuerUrl"`
	RequiredClaims             map[string]string `json:"requiredClaims"`
	UsernameClaim              string            `json:"usernameClaim"`
	UsernamePrefix             string            `json:"usernamePrefix"`
}

func (AssociateResource) Uri() string {
	return "/eks/identity-provider-config/associate"
}
func (AssociateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(IdentityProviderConfigAssociateInput{}, eks.AssociateIdentityProviderConfigOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
