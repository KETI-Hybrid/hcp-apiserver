package snapshot

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type CreateOrUpdateResource struct {
	docs.GetNotSupported
	docs.PostNotSupported
	docs.DeleteNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type CreateOrUpdateSnapshot struct {
	ResourceGroupName string                        `json:"resourceGroupName"`
	ClusterName       string                        `json:"clusterName"`
	SnapshotProperty  *armcontainerservice.Snapshot `json:"snapshotProperty"`
}

func (CreateOrUpdateResource) Uri() string {
	return "/aks/snapshot/createorupdate"
}
func (CreateOrUpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(CreateOrUpdateSnapshot{}, armcontainerservice.SnapshotsClientCreateOrUpdateResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
