package resolveprivatelink

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
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
	request, response := util.DocWithReq(POSTResolvePrivateLinkServiceId{}, armcontainerservice.ResolvePrivateLinkServiceIDClientPOSTResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
