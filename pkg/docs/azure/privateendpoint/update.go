package privateendpoint

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type UpdateResource struct {
	docs.GetNotSupported
	docs.PostNotSupported
	docs.DeleteNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type UpdatePrivateEndpointConnection struct {
	ResourceGroupName             string                                         `json:"resourceGroupName"`
	ClusterName                   string                                         `json:"clusterName"`
	PrivateEndpointConnectionName string                                         `json:"privateEndpointConnectionName"`
	PrivateEndpointConnection     *armcontainerservice.PrivateEndpointConnection `json:"privateEndpointConnectionProperty"`
}

func (UpdateResource) Uri() string {
	return "/aks/privateEndpointConnection/update"
}
func (UpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &UpdatePrivateEndpointConnection{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := PrivateEndpointConnectionsClient.Update(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, inputRequest.PrivateEndpointConnectionName, *inputRequest.PrivateEndpointConnection, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
