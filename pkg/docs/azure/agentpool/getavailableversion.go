package agentpool

import (
	"fmt"
	"hcp-apiserver/pkg/apis"
	"net/http"
	"reflect"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type GetAvailableAgentPoolVersionsResource struct {
	apis.DeleteNotSupported
	apis.PostNotSupported
	apis.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type GetAvailableAgentPoolVersions struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
}

func (GetAvailableAgentPoolVersionsResource) Uri() string {
	return "/aks/agentpool/availableAgentPoolVersions"
}
func (GetAvailableAgentPoolVersionsResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	t := reflect.TypeOf(GetAvailableAgentPoolVersions{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s: %s\n", field.Name, field.Type)
	}

	typeMap := make(map[string]string)

	t2 := reflect.TypeOf(armcontainerservice.AgentPoolsClientGetAvailableAgentPoolVersionsResponse{})
	for i := 0; i < t2.NumField(); i++ {
		field := t2.Field(i)
		typeMap[field.Name] = field.Type.String()
	}
	return apis.Response{Code: 200, Data: typeMap}

}
