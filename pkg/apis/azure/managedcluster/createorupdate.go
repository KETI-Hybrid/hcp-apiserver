package managedcluster

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
type CreateOrUpdateManagedCluster struct {
	ResourceGroupName string                              `json:"resourceGroupName"`
	ClusterName       string                              `json:"clusterName"`
	ClusterProperty   *armcontainerservice.ManagedCluster `json:"clusterProperty"`
}

func (CreateOrUpdateResource) Uri() string {
	return "/aks/managedClusters/createorupdate"
}
func (CreateOrUpdateResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &CreateOrUpdateManagedCluster{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := inputRequest.ClusterProperty
	ctx := context.Background()
	result, err := ManagedClustersClient.BeginCreateOrUpdate(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, *realInput, nil)
	if err != nil {
		klog.Errorln(err)
	}
	resp, err := result.PollUntilDone(ctx, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
