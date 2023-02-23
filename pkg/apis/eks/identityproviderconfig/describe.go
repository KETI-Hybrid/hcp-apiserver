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

type DescribeResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
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
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &IdentityProviderConfigDescribeInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DescribeIdentityProviderConfigInput{
		ClusterName: aws.String(inputReq.ClusterName),
		IdentityProviderConfig: &eks.IdentityProviderConfig{
			Name: aws.String(inputReq.IdentityProviderConfig.Name),
			Type: aws.String(inputReq.IdentityProviderConfig.Type),
		},
	}
	result, err := IdentityProviderConfigClient.DescribeIdentityProviderConfig(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
