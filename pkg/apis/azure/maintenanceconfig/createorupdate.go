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

type CreateOrUpdateResource struct {
	apis.GetNotSupported
	apis.PostNotSupported
	apis.DeleteNotSupported
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
func (CreateOrUpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &CreateOrUpdateMaintenanceConfiguration{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := inputRequest.ConfigProperty
	ctx := context.Background()
	result, err := MaintenanceConfigurationsClient.CreateOrUpdate(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, inputRequest.ConfigName, *realInput, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
