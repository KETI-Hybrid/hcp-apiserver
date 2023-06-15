package resource

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

type UntagResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type ResourceUntagtInput struct {
	ResourceArn string   `json:"resourceArn"`
	TagKeys     []string `json:"tagKeys"`
}

func (UntagResource) Uri() string {
	return "/eks/resource/untag"
}
func (UntagResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ResourceUntagtInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.UntagResourceInput{
		ResourceArn: aws.String(inputReq.ResourceArn),
		TagKeys:     aws.StringSlice(inputReq.TagKeys),
	}

	result, err := ResourceClient.UntagResource(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
