package cluster

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

type RegisterResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}
type ClusterRegisterInput struct {
	ClientRequestToken string                 `json:"clientRequestToken"`
	ConnectorConfig    ConnectorConfigRequest `json:"connectorConfig"`
	Name               string                 `json:"name"`
	Tags               map[string]string      `json:"tags"`
}
type ConnectorConfigRequest struct {
	Provider string `json:"provider"`
	RoleArn  string `json:"roleArn"`
}

func (RegisterResource) Uri() string {
	return "/eks/cluster/register"
}
func (RegisterResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ClusterRegisterInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.RegisterClusterInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		ConnectorConfig: &eks.ConnectorConfigRequest{
			Provider: aws.String(inputReq.ConnectorConfig.Provider),
			RoleArn:  aws.String(inputReq.ConnectorConfig.RoleArn),
		},
		Name: aws.String(inputReq.Name),
		Tags: aws.StringMap(inputReq.Tags),
	}

	result, err := ClusterClient.RegisterCluster(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
