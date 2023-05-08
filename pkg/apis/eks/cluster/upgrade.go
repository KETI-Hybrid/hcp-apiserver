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

type UpgradeResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type ClusterUpgradeInput struct {
	ClientRequestToken string `json:"clientRequestToken"`
	Name               string `json:"name"`
	Version            string `json:"version"`
}

func (UpgradeResource) Uri() string {
	return "/eks/cluster/upgrade"
}
func (UpgradeResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ClusterUpgradeInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.UpdateClusterVersionInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		Name:               aws.String(inputReq.Name),
		Version:            aws.String(inputReq.Version),
	}

	result, err := ClusterClient.UpdateClusterVersion(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
