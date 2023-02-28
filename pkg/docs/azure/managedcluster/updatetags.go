package managedcluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type UpdateTagsResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type UpdateTags struct {
	ResourceGroupName string                          `json:"resourceGroupName"`
	ClusterName       string                          `json:"clusterName"`
	Tags              *armcontainerservice.TagsObject `json:"tags"`
}

func (UpdateTagsResource) Uri() string {
	return "/aks/managedClusters/updateTags"
}
func (UpdateTagsResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(UpdateTags{}, armcontainerservice.ManagedClustersClientUpdateTagsResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
