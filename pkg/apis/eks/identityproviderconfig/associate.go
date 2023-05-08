package identityproviderconfig

import (
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type AssociateResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
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
func (AssociateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &IdentityProviderConfigAssociateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.AssociateIdentityProviderConfigInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		ClusterName:        aws.String(inputReq.ClusterName),
		Oidc: &eks.OidcIdentityProviderConfigRequest{
			ClientId:                   aws.String(inputReq.Oidc.ClientId),
			GroupsClaim:                aws.String(inputReq.Oidc.GroupsClaim),
			GroupsPrefix:               aws.String(inputReq.Oidc.GroupsPrefix),
			IdentityProviderConfigName: aws.String(inputReq.Oidc.IdentityProviderConfigName),
			IssuerUrl:                  aws.String(inputReq.Oidc.IssuerUrl),
			RequiredClaims:             aws.StringMap(inputReq.Oidc.RequiredClaims),
			UsernameClaim:              aws.String(inputReq.Oidc.UsernameClaim),
			UsernamePrefix:             aws.String(inputReq.Oidc.UsernamePrefix),
		},
		Tags: aws.StringMap(inputReq.Tags),
	}

	result, err := IdentityProviderConfigClient.AssociateIdentityProviderConfig(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
