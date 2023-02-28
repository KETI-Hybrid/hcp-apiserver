package encryptionconfig

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type AssociateResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type EncryptionConfigAssociateInput struct {
	ClientRequestToken string             `json:"clientRequestToken"`
	ClusterName        string             `json:"clusterName"`
	EncryptionConfig   []EncryptionConfig `json:"encryptionConfig"`
}
type EncryptionConfig struct {
	Provider  Provider `json:"provider"`
	Resources []string `json:"resources"`
}
type Provider struct {
	KeyArn string
}

func (AssociateResource) Uri() string {
	return "/eks/encryption-config/associate"
}
func (AssociateResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(EncryptionConfigAssociateInput{}, eks.AssociateEncryptionConfigOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
