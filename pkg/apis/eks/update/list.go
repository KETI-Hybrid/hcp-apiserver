package update

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

type ListResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type UpdateListInput struct {
	AddonName     string `json:"addonName"`
	MaxResults    int64  `json:"maxResults"`
	Name          string `json:"name"`
	NextToken     string `json:"nextToken"`
	NodegroupName string `json:"nodegroupName"`
}

func (ListResource) Uri() string {
	return "/eks/update/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &UpdateListInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.ListUpdatesInput{
		AddonName:     aws.String(inputReq.AddonName),
		MaxResults:    aws.Int64(inputReq.MaxResults),
		Name:          aws.String(inputReq.Name),
		NextToken:     aws.String(inputReq.NextToken),
		NodegroupName: aws.String(inputReq.NodegroupName),
	}

	result, err := UpdateClient.ListUpdates(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
