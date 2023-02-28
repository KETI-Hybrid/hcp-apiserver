package addon

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

type ListResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type AddonListInput struct {
	ClusterName string `json:"clusterName"`
	MaxResults  int64  `json:"maxResults"`
	NextToken   string `json:"nextToken"`
}

func (ListResource) Uri() string {
	return "/eks/addon/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &AddonListInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.ListAddonsInput{
		ClusterName: aws.String(inputReq.ClusterName),
		MaxResults:  aws.Int64(inputReq.MaxResults),
		NextToken:   aws.String(inputReq.NextToken),
	}

	result, err := AddonClient.ListAddons(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
