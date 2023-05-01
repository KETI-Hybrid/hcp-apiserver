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

type DeleteResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.PostNotSupported
}

type AddonDeleteInput struct {
	AddonName   string `json:"addonName"`
	ClusterName string `json:"clusterName"`
	Preserve    bool   `json:"preserve"`
}

func (DeleteResource) Uri() string {
	return "/eks/addon/delete"
}
func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &AddonDeleteInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DeleteAddonInput{
		AddonName:   aws.String(inputReq.AddonName),
		ClusterName: aws.String(inputReq.ClusterName),
		Preserve:    aws.Bool(inputReq.Preserve),
	}

	result, err := AddonClient.DeleteAddon(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
