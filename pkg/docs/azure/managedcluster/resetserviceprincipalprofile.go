package managedcluster

import (
	"fmt"
	"hcp-apiserver/pkg/docs"
	"net/http"
	"reflect"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type ResetServicePrincipalProfileResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type ResetServicePrincipalProfile struct {
	ResourceGroupName        string                                                     `json:"resourceGroupName"`
	ClusterName              string                                                     `json:"clusterName"`
	ServicePrincipalProperty *armcontainerservice.ManagedClusterServicePrincipalProfile `json:"servicePrincipalProperty"`
}

func (ResetServicePrincipalProfileResource) Uri() string {
	return "/aks/managedClusters/resetServicePrincipalProfile"
}
func (ResetServicePrincipalProfileResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	t := reflect.TypeOf(ResetServicePrincipalProfile{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s: %s\n", field.Name, field.Type)
	}

	typeMap := make(map[string]string)

	t2 := reflect.TypeOf(armcontainerservice.ManagedClustersClientResetServicePrincipalProfileResponse{})
	for i := 0; i < t2.NumField(); i++ {
		field := t2.Field(i)
		typeMap[field.Name] = field.Type.String()
	}
	return docs.Response{Code: 200, Data: typeMap}

}
