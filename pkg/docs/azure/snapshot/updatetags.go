package snapshot

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type UpdateTagsResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.GetNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type UpdateTagSnapshot struct {
	ResourceGroupName string                          `json:"resourceGroupName"`
	ClusterName       string                          `json:"clusterName"`
	Tags              *armcontainerservice.TagsObject `json:"tags"`
}

func (UpdateTagsResource) Uri() string {
	return "/aks/snapshot/updateTag"
}
func (UpdateTagsResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(UpdateTagSnapshot{}, armcontainerservice.SnapshotsClientUpdateTagsResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
