package fargateprofile

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
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
	request, response := util.DocWithReq(FargateProfileDescribeInput{}, eks.DescribeFargateProfileOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
