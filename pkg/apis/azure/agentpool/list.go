package agentpool

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

type ListResource struct {
	apis.DeleteNotSupported
	apis.PostNotSupported
	apis.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type ListAgentPool struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
}

func (ListResource) Uri() string {
	return "/aks/agentpool/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &ListAgentPool{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	result := make([]armcontainerservice.AgentPoolsClientListResponse, 0)
	items := AgentPoolsClient.NewListPager(inputRequest.ResourceGroupName, inputRequest.ClusterName, nil)
	ctx := context.Background()

	for items.More() {
		resp, err := items.NextPage(ctx)
		if err != nil {
			klog.Errorln(err)
		}
		result = append(result, resp)
	}

	return apis.Response{Code: 200, Data: result}
}
