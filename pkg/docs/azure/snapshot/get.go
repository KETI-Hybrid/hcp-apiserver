package snapshot

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type GetResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type GetSnapshot struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
}

func (GetResource) Uri() string {
	return "/aks/snapshot/get"
}
func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &DeleteSnapshot{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := SnapshotsClient.Get(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
