package workernode

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type GetResource struct {
	docs.PutNotSupported
	docs.PostNotSupported
	docs.DeleteNotSupported
}

type Get struct {
	ClusterUUID string `json:"clusteruuid"`
}

func (GetResource) Uri() string {
	return "/nks/workerNode/get"
}

func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	client := types.GetNKSClient()
	containerService := client.Client.V2Api
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Get{}
	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	resp, err := containerService.ClustersUuidNodesGet(ctx, &inputRequest.ClusterUUID)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: resp}
}
