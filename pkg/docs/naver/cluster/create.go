package cluster

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type CreateResource struct {
	docs.PutNotSupported
	docs.DeleteNotSupported
	docs.GetNotSupported
}

type Create struct {
	Cluster *vnks.ClusterInputBody `json:"cluster"`
}

func (CreateResource) Uri() string {
	return "/nks/cluster/create"
}

func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	client := types.GetNKSClient()
	containerService := client.Client.V2Api
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Create{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	resp, err := containerService.ClustersPost(ctx, inputRequest.Cluster)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: resp}
}
