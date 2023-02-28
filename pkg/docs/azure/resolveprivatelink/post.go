package resolveprivatelink

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

type POSTResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type POSTResolvePrivateLinkServiceId struct {
	ResourceGroupName   string                                   `json:"resourceGroupName"`
	ClusterName         string                                   `json:"clusterName"`
	PrivateLinkResource *armcontainerservice.PrivateLinkResource `json:"privateLinkProperty"`
}

func (POSTResource) Uri() string {
	return "/aks/resolvePrivateLinkServiceId/post"
}
func (POSTResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &POSTResolvePrivateLinkServiceId{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := ResolvePrivateLinkServiceIdClient.POST(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, *inputRequest.PrivateLinkResource, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
