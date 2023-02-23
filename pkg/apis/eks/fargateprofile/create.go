package fargateprofile

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

type FargateProfileCreateInput struct {
	ClientRequestToken  string                   `json:"clientRequestToken"`
	ClusterName         string                   `json:"clusterName"`
	FargateProfileName  string                   `json:"fargateProfileName"`
	PodExecutionRoleArn string                   `json:"podExecutionRoleArn"`
	Selectors           []FargateProfileSelector `json:"selectors"`
	Subnets             []string                 `json:"subnets"`
	Tags                map[string]string        `json:"tags"`
}
type FargateProfileSelector struct {
	Labels    map[string]string `json:"labels"`
	Namespace string            `json:"namespace"`
}

func (CreateResource) Uri() string {
	return "/eks/fargate-profile/create"
}
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &FargateProfileCreateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.CreateFargateProfileInput{
		ClientRequestToken:  aws.String(inputReq.ClientRequestToken),
		ClusterName:         aws.String(inputReq.ClusterName),
		FargateProfileName:  aws.String(inputReq.FargateProfileName),
		PodExecutionRoleArn: aws.String(inputReq.PodExecutionRoleArn),
		Selectors:           make([]*eks.FargateProfileSelector, 0),
		Subnets:             aws.StringSlice(inputReq.Subnets),
		Tags:                aws.StringMap(inputReq.Tags),
	}
	for _, seletor := range inputReq.Selectors {
		realInput.Selectors = append(realInput.Selectors, &eks.FargateProfileSelector{
			Labels:    aws.StringMap(seletor.Labels),
			Namespace: aws.String(seletor.Namespace),
		})
	}

	result, err := FargateProfileClient.CreateFargateProfile(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
