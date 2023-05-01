package addon

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

type UpdateResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type AddonUpdateInput struct {
	AddonName             string `json:"addonName"`
	AddonVersion          string `json:"addonVersion"`
	ClientRequestToken    string `json:"clientRequestToken"`
	ClusterName           string `json:"clusterName"`
	ConfigurationValues   string `json:"configurationValues"`
	ResolveConflicts      string `json:"resolveConflicts"`
	ServiceAccountRoleArn string `json:"serviceAccountRoleArn"`
}

func (UpdateResource) Uri() string {
	return "/eks/addon/update"
}
func (UpdateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &AddonUpdateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.UpdateAddonInput{
		AddonName:             aws.String(inputReq.AddonName),
		AddonVersion:          aws.String(inputReq.AddonVersion),
		ClientRequestToken:    aws.String(inputReq.ClientRequestToken),
		ClusterName:           aws.String(inputReq.ClusterName),
		ConfigurationValues:   aws.String(inputReq.ConfigurationValues),
		ResolveConflicts:      aws.String(inputReq.ResolveConflicts),
		ServiceAccountRoleArn: aws.String(inputReq.ServiceAccountRoleArn),
	}

	result, err := AddonClient.UpdateAddon(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
