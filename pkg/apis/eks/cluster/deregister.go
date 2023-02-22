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

type DeregisterResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type ClusterDeregisterInput struct {
	Name string `json:"name"`
}

func (DeregisterResource) Uri() string {
	return "/eks/cluster/deregister"
}
func (DeregisterResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ClusterDeregisterInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DeregisterClusterInput{
		Name: aws.String(inputReq.Name),
	}

	result, err := ClusterClient.DeregisterCluster(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
