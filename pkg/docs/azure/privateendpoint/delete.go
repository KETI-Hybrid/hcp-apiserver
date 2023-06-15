package privateendpoint

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type DeleteResource struct {
	docs.GetNotSupported
	docs.PostNotSupported
	docs.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type DeletePrivateEndpointConnection struct {
	ResourceGroupName             string `json:"resourceGroupName"`
	ClusterName                   string `json:"clusterName"`
	PrivateEndpointConnectionName string `json:"privateEndpointConnectionName"`
}

func (DeleteResource) Uri() string {
	return "/aks/privateEndpointConnection/delete"
}
func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &DeletePrivateEndpointConnection{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := PrivateEndpointConnectionsClient.BeginDelete(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, inputRequest.PrivateEndpointConnectionName, nil)
	if err != nil {
		klog.Errorln(err)
	}
	resp, err := result.PollUntilDone(ctx, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: resp}
}
