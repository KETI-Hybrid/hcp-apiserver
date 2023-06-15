package fargateprofile

import (
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type DeleteResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.GetNotSupported
}

type FargateProfileDeleteInput struct {
	ClusterName        string `json:"clusterName"`
	FargateProfileName string `json:"fargateProfileName"`
}

func (DeleteResource) Uri() string {
	return "/eks/fargate-profile/delete"
}
func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &FargateProfileDeleteInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DeleteFargateProfileInput{
		ClusterName:        aws.String(inputReq.ClusterName),
		FargateProfileName: aws.String(inputReq.FargateProfileName),
	}

	result, err := FargateProfileClient.DeleteFargateProfile(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
