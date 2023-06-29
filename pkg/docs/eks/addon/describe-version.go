package addon

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type DescribeVersionResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type AddonDescribeVersionsInput struct {
	AddonName         string   `json:"addonName"`
	KubernetesVersion string   `json:"kubernetesVersion"`
	MaxResults        int64    `json:"maxResults"`
	NextToken         string   `json:"nextToken"`
	Owners            []string `json:"owners"`
	Publishers        []string `json:"publishers"`
	Types             []string `json:"types"`
}

func (DescribeVersionResource) Uri() string {
	return "/eks/addon/describe-version"
}
func (DescribeVersionResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(AddonDescribeVersionsInput{}, eks.DescribeAddonVersionsOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
