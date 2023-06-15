package nodegroup

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

type NodeGroupDescribeInput struct {
	ClusterName   string `json:"clusterName"`
	NodegroupName string `json:"nodegroupName"`
}

func (DescribeResource) Uri() string {
	return "/eks/nodegroup/describe"
}
func (DescribeResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &NodeGroupDescribeInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DescribeNodegroupInput{
		ClusterName:   aws.String(inputReq.ClusterName),
		NodegroupName: aws.String(inputReq.NodegroupName),
	}

	result, err := NodeGroupClient.DescribeNodegroup(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
