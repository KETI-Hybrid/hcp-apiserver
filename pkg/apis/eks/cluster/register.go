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
	Provider *string `locationName:"provider"`
	RoleArn  *string `locationName:"roleArn"`
}

func (RegisterResource) Uri() string {
	return "/eks/cluster/register"
}
func (RegisterResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ClusterListInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.ListClustersInput{
		Include:    aws.StringSlice(inputReq.Include),
		MaxResults: aws.Int64(inputReq.MaxResults),
		NextToken:  aws.String(inputReq.NextToken),
	}

	result, err := ClusterClient.ListClusters(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
