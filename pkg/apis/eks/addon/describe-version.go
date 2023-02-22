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

type DescribeVersionResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
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
func (DescribeVersionResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &AddonDescribeVersionsInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DescribeAddonVersionsInput{
		AddonName:         aws.String(inputReq.AddonName),
		KubernetesVersion: aws.String(inputReq.KubernetesVersion),
		MaxResults:        aws.Int64(inputReq.MaxResults),
		NextToken:         aws.String(inputReq.NextToken),
		Owners:            aws.StringSlice(inputReq.Owners),
		Publishers:        aws.StringSlice(inputReq.Publishers),
		Types:             aws.StringSlice(inputReq.Types),
	}

	result, err := AddonClient.DescribeAddonVersions(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
