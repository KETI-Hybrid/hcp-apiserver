package nodegroup

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

type NodeGroupUpgradeInput struct {
	ClientRequestToken string                      `json:"clientRequestToken"`
	ClusterName        string                      `json:"clusterName"`
	Force              bool                        `json:"force"`
	LaunchTemplate     LaunchTemplateSpecification `json:"launchTemplate"`
	NodegroupName      string                      `json:"nodegroupName"`
	ReleaseVersion     string                      `json:"releaseVersion"`
	Version            string                      `json:"version"`
}

func (UpgradeResource) Uri() string {
	return "/eks/nodegroup/upgrade"
}
func (UpgradeResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &NodeGroupUpgradeInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.UpdateNodegroupVersionInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		ClusterName:        aws.String(inputReq.ClusterName),
		Force:              aws.Bool(inputReq.Force),
		LaunchTemplate: &eks.LaunchTemplateSpecification{
			Id:      aws.String(inputReq.LaunchTemplate.Id),
			Name:    aws.String(inputReq.LaunchTemplate.Name),
			Version: aws.String(inputReq.LaunchTemplate.Version),
		},
		NodegroupName:  aws.String(inputReq.NodegroupName),
		ReleaseVersion: aws.String(inputReq.ReleaseVersion),
		Version:        aws.String(inputReq.Version),
	}

	result, err := NodeGroupClient.UpdateNodegroupVersion(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
