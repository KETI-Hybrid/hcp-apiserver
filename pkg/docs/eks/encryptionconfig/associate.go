package encryptionconfig

import (
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &EncryptionConfigAssociateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.AssociateEncryptionConfigInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		ClusterName:        aws.String(inputReq.ClusterName),
		EncryptionConfig:   make([]*eks.EncryptionConfig, 0),
	}
	for _, encryptionConfig := range inputReq.EncryptionConfig {
		realInput.EncryptionConfig = append(realInput.EncryptionConfig, &eks.EncryptionConfig{
			Provider: &eks.Provider{
				KeyArn: aws.String(encryptionConfig.Provider.KeyArn),
			},
			Resources: aws.StringSlice(encryptionConfig.Resources),
		})
	}

	result, err := EncryptionConfigClient.AssociateEncryptionConfig(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
