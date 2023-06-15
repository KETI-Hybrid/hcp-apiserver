package nodepools

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type DeleteResource struct {
	apis.PutNotSupported
	apis.PostNotSupported
	apis.GetNotSupported
}

type Delete struct {
	ClusterUUID string `json:"clusteruuid"`
	NodeNumber  string `json:"nodeNumber"`
}

type DeleteResp struct {
	ClusterUUID string `json:"clusteruuid"`
	NodeNumber  string `json:"nodeNumber"`
	Status      bool   `json:"status"`
}

func (DeleteResource) Uri() string {
	return "/nks/nodepool/delete"
}

func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetNKSClient()
	containerService := client.Client.V2Api
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Delete{}
	resp := &DeleteResp{}
	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	err = containerService.ClustersUuidNodePoolInstanceNoDelete(ctx, &inputRequest.ClusterUUID, &inputRequest.NodeNumber)
	if err != nil {
		klog.Errorln(err)
	}
	resp.ClusterUUID = inputRequest.ClusterUUID
	resp.NodeNumber = inputRequest.NodeNumber
	resp.Status = true
	return apis.Response{Code: 200, Data: resp}
}
