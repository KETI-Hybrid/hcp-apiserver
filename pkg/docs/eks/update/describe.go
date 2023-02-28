package update

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

type UpdateDescribeInput struct {
	AddonName     string `json:"addonName"`
	Name          string `json:"name"`
	NodegroupName string `json:"nodegroupName"`
	UpdateId      string `json:"updateId"`
}

func (DescribeResource) Uri() string {
	return "/eks/update/describe"
}
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &UpdateDescribeInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DescribeUpdateInput{
		AddonName:     aws.String(inputReq.AddonName),
		Name:          aws.String(inputReq.Name),
		NodegroupName: aws.String(inputReq.NodegroupName),
		UpdateId:      aws.String(inputReq.UpdateId),
	}

	result, err := UpdateClient.DescribeUpdate(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
