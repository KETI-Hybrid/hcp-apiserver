package maintenanceconfig

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
type CreateOrUpdateMaintenanceConfiguration struct {
	ResourceGroupName string                                        `json:"resourceGroupName"`
	ClusterName       string                                        `json:"clusterName"`
	ConfigName        string                                        `json:"configName"`
	ConfigProperty    *armcontainerservice.MaintenanceConfiguration `json:"configProperty"`
}

func (CreateOrUpdateResource) Uri() string {
	return "/aks/maintenanceConfigurations/createorupdate"
}
func (CreateOrUpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(CreateOrUpdateMaintenanceConfiguration{}, armcontainerservice.MaintenanceConfigurationsClientCreateOrUpdateResponse{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
