package managedcluster

import (
	"hcp-apiserver/pkg/docs"
	"net/http"
	"reflect"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

type ListResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.PutNotSupported
}

func (ListResource) Uri() string {
	return "/aks/managedClusters/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	typeMap := make(map[string]string)

	t2 := reflect.TypeOf(armcontainerservice.ManagedClustersClientListResponse{})
	for i := 0; i < t2.NumField(); i++ {
		field := t2.Field(i)
		typeMap[field.Name] = field.Type.String()
	}
	return docs.Response{Code: 200, Data: typeMap}

}
