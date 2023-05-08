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

type DisassociateResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type IdentityProviderConfigDisassociateInput struct {
	ClientRequestToken     string                 `json:"clientRequestToken"`
	ClusterName            string                 `json:"clusterName"`
	IdentityProviderConfig IdentityProviderConfig `json:"identityProviderConfig"`
}

func (DisassociateResource) Uri() string {
	return "/eks/identity-provider-config/disassociate"
}
func (DisassociateResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &IdentityProviderConfigDisassociateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DisassociateIdentityProviderConfigInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		ClusterName:        aws.String(inputReq.ClusterName),
		IdentityProviderConfig: &eks.IdentityProviderConfig{
			Name: aws.String(inputReq.IdentityProviderConfig.Name),
			Type: aws.String(inputReq.IdentityProviderConfig.Type),
		},
	}

	result, err := IdentityProviderConfigClient.DisassociateIdentityProviderConfig(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
