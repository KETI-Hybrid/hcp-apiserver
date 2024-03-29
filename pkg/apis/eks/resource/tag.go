package resource

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

type TagResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type ResourceTagInput struct {
	ResourceArn string            `json:"resourceArn"`
	Tags        map[string]string `json:"tags"`
}

func (TagResource) Uri() string {
	return "/eks/resource/tag"
}
func (TagResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ResourceTagInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.TagResourceInput{
		ResourceArn: aws.String(inputReq.ResourceArn),
		Tags:        aws.StringMap(inputReq.Tags),
	}

	result, err := ResourceClient.TagResource(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
