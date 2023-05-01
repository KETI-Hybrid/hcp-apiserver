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

type DescribeConfigResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type AddonDescribeConfigInput struct {
	AddonName    string `json:"addonName"`
	AddonVersion string `json:"addonVersion"`
}

func (DescribeConfigResource) Uri() string {
	return "/eks/addon/describe-confog"
}
func (DescribeConfigResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &AddonDescribeConfigInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DescribeAddonConfigurationInput{
		AddonName:    aws.String(inputReq.AddonName),
		AddonVersion: aws.String(inputReq.AddonVersion),
	}

	result, err := AddonClient.DescribeAddonConfiguration(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
