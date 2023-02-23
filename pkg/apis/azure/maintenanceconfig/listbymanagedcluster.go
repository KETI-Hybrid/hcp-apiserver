package maintenanceconfig

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"io/ioutil"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListByManagedClusterResource struct {
	apis.PutNotSupported
	apis.PostNotSupported
	apis.DeleteNotSupported
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
func (ListByManagedClusterResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &CreateOrUpdateMaintenanceConfiguration{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result := make([]armcontainerservice.MaintenanceConfigurationsClientListByManagedClusterResponse, 0)
	items := MaintenanceConfigurationsClient.NewListByManagedClusterPager(inputRequest.ResourceGroupName, inputRequest.ClusterName, nil)
	for items.More() {
		tmp, err := items.NextPage(ctx)
		if err != nil {
			klog.Errorln(err)
		}

		result = append(result, tmp)
	}
	return apis.Response{Code: 200, Data: result}
}
