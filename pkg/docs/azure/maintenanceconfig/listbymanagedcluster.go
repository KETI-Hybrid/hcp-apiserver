package maintenanceconfig

import (
	"fmt"
	"hcp-apiserver/pkg/docs"
	"net/http"
	"reflect"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type ListByManagedClusterResource struct {
	docs.PutNotSupported
	docs.PostNotSupported
	docs.DeleteNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type ListByManagedCluster struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
}

func (ListByManagedClusterResource) Uri() string {
	return "/aks/maintenanceConfigurations/list"
}
func (ListByManagedClusterResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	t := reflect.TypeOf(ListByManagedCluster{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s: %s\n", field.Name, field.Type)
	}

	typeMap := make(map[string]string)

	t2 := reflect.TypeOf(armcontainerservice.MaintenanceConfigurationsClientListByManagedClusterResponse{})
	for i := 0; i < t2.NumField(); i++ {
		field := t2.Field(i)
		typeMap[field.Name] = field.Type.String()
	}
	return docs.Response{Code: 200, Data: typeMap}

}
