package agentpool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type GetResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type GetAgentPool struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
	AgentPoolName     string `json:"agentPoolName"`
}

func (GetResource) Uri() string {
	return "/aks/agentpool/get"
}
func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(GetAgentPool{}, armcontainerservice.AgentPoolsClientGetResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}

}
