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

type CreateResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type AddonCreateInput struct {
	AddonName             string            `json:"addonName"`
	AddonVersion          string            `json:"addonVersion"`
	ClientRequestToken    string            `json:"clientRequestToken"`
	ClusterName           string            `json:"clusterName"`
	ConfigurationValues   string            `json:"configurationValues"`
	ResolveConflicts      string            `json:"resolveConflicts"`
	ServiceAccountRoleArn string            `json:"serviceAccountRoleArn"`
	Tags                  map[string]string `json:"tags"`
}

func (CreateResource) Uri() string {
	return "/eks/addon/create"
}
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &AddonCreateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.CreateAddonInput{
		AddonName:             aws.String(inputReq.AddonName),
		AddonVersion:          aws.String(inputReq.AddonVersion),
		ClusterName:           aws.String(inputReq.ClusterName),
		ServiceAccountRoleArn: aws.String(inputReq.ServiceAccountRoleArn),
		ClientRequestToken:    aws.String(inputReq.ClientRequestToken),
		ConfigurationValues:   aws.String(inputReq.ConfigurationValues),
		ResolveConflicts:      aws.String(inputReq.ResolveConflicts),
		Tags:                  aws.StringMap(inputReq.Tags),
	}

	result, err := AddonClient.CreateAddon(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
