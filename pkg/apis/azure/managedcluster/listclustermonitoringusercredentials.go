package managedcluster

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListClusterMonitoringUserCredentialsResource struct {
	apis.DeleteNotSupported
	apis.GetNotSupported
	apis.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type ListClusterMonitoringUserCredentials struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
}

func (ListClusterMonitoringUserCredentialsResource) Uri() string {
	return "/aks/managedClusters/listClusterMonitoringUserCredentials"
}
func (ListClusterMonitoringUserCredentialsResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &ListClusterMonitoringUserCredentials{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := ManagedClustersClient.ListClusterMonitoringUserCredentials(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
