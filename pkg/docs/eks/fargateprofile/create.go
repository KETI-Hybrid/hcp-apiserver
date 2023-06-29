package fargateprofile

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type CreateResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
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
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(FargateProfileCreateInput{}, eks.CreateFargateProfileOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
