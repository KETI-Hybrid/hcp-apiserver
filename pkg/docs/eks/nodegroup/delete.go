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

type DeleteResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.GetNotSupported
}

type NodeGroupDeleteInput struct {
	ClusterName   string `json:"clusterName"`
	NodegroupName string `json:"nodegroupName"`
}

func (DeleteResource) Uri() string {
	return "/eks/nodegroup/delete"
}
func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &NodeGroupDeleteInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.DeleteNodegroupInput{
		ClusterName:   aws.String(inputReq.ClusterName),
		NodegroupName: aws.String(inputReq.NodegroupName),
	}

	result, err := NodeGroupClient.DeleteNodegroup(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
