package nodepools

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

type UpdateResource struct {
	docs.GetNotSupported
	docs.DeleteNotSupported
	docs.PostNotSupported
}

type Update struct {
	ClusterUUID string                   `json:"clusteruuid"`
	NodeNumber  string                   `json:"nodeNumber"`
	UpdateBody  *vnks.NodePoolUpdateBody `json:"updateBody"`
}

type UpdateResp struct {
	ClusterUUID string `json:"clusteruuid"`
	NodeNumber  string `json:"nodeNumber"`
	Status      bool   `json:"status"`
}

func (UpdateResource) Uri() string {
	return "/nks/nodepool/update"
}

func (UpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	client := types.GetNKSClient()
	containerService := client.Client.V2Api
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Update{}
	resp := &UpdateResp{}
	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	err = containerService.ClustersUuidNodePoolInstanceNoPatch(ctx, inputRequest.UpdateBody, &inputRequest.ClusterUUID, &inputRequest.NodeNumber)
	if err != nil {
		klog.Errorln(err)
	}
	resp.ClusterUUID = inputRequest.ClusterUUID
	resp.NodeNumber = inputRequest.NodeNumber
	resp.Status = true
	return docs.Response{Code: 200, Data: resp}
}
