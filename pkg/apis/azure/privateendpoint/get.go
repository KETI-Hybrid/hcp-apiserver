package privateendpoint

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type GetResource struct {
	apis.DeleteNotSupported
	apis.PostNotSupported
	apis.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type GetPrivateEndpointConnection struct {
	ResourceGroupName             string `json:"resourceGroupName"`
	ClusterName                   string `json:"clusterName"`
	PrivateEndpointConnectionName string `json:"privateEndpointConnectionName"`
}

func (GetResource) Uri() string {
	return "/aks/privateEndpointConnection/get"
}
func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &GetPrivateEndpointConnection{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := PrivateEndpointConnectionsClient.Get(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, inputRequest.PrivateEndpointConnectionName, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
