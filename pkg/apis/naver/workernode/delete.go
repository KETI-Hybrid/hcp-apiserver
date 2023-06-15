package workernode

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
	NodePoolID  string `json:"nodePoolID"`
	NodeNumber  string `json:"nodeNumber"`
}

type DeleteResp struct {
	ClusterUUID string `json:"clusteruuid"`
	NodePoolID  string `json:"nodePoolID"`
	NodeNumber  string `json:"nodeNumber"`
	Status      bool   `json:"status"`
}

func (DeleteResource) Uri() string {
	return "/nks/workerNode/delete"
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
	err = containerService.ClustersUuidNodesInstanceNoDelete(ctx, &inputRequest.ClusterUUID, &inputRequest.NodeNumber, &inputRequest.NodePoolID)
	if err != nil {
		klog.Errorln(err)
	}
	resp.ClusterUUID = inputRequest.ClusterUUID
	resp.NodeNumber = inputRequest.NodeNumber
	resp.NodePoolID = inputRequest.NodePoolID
	resp.Status = true
	return apis.Response{Code: 200, Data: resp}
}
