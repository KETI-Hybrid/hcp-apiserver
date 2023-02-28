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

type DescribeResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type FargateProfileDescribeInput struct {
	ClusterName        string `json:"clusterName"`
	FargateProfileName string `json:"fargateProfileName"`
}

func (DescribeResource) Uri() string {
	return "/eks/fargate-profile/describe"
}
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &FargateProfileDescribeInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DescribeFargateProfileInput{
		ClusterName:        aws.String(inputReq.ClusterName),
		FargateProfileName: aws.String(inputReq.FargateProfileName),
	}

	result, err := FargateProfileClient.DescribeFargateProfile(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
