package nodepools

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type CreateResource struct {
	apis.PutNotSupported
	apis.DeleteNotSupported
	apis.GetNotSupported
}

type Create struct {
	ClusterUUID string                     `json:"clusteruuid"`
	NodePool    *vnks.NodePoolCreationBody `json:"nodepool"`
}

type CreateResp struct {
	ClusterUUID  string `json:"clusteruuid"`
	NodePoolName string `json:"nodepoolName"`
	Status       bool   `json:"status"`
}

func (CreateResource) Uri() string {
	return "/nks/nodepool/create"
}

func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetNKSClient()
	containerService := client.Client.V2Api
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Create{}

	resp := &CreateResp{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	err = containerService.ClustersUuidNodePoolPost(ctx, inputRequest.NodePool, &inputRequest.ClusterUUID)
	if err != nil {
		klog.Errorln(err)
	}

	resp.ClusterUUID = inputRequest.ClusterUUID
	resp.NodePoolName = *inputRequest.NodePool.Name
	resp.Status = true
	return apis.Response{Code: 200, Data: resp}
}
