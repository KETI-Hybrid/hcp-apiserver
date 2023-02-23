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

type RunCommandResource struct {
	apis.DeleteNotSupported
	apis.GetNotSupported
	apis.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type RunCommand struct {
	ResourceGroupName string                                 `json:"resourceGroupName"`
	ClusterName       string                                 `json:"clusterName"`
	RunCommandRequest *armcontainerservice.RunCommandRequest `json:"runCommandRequest"`
}

func (RunCommandResource) Uri() string {
	return "/aks/managedClusters/runCommand"
}
func (RunCommandResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &RunCommand{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := inputRequest.RunCommandRequest
	ctx := context.Background()
	result, err := ManagedClustersClient.BeginRunCommand(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, *realInput, nil)
	if err != nil {
		klog.Errorln(err)
	}
	resp, err := result.PollUntilDone(ctx, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
