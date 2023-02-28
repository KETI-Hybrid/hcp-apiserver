package snapshot

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

type ListByResourceGroupResource struct {
	apis.DeleteNotSupported
	apis.PostNotSupported
	apis.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type ListByResourceGroupSnapshot struct {
	ResourceGroupName string `json:"resourceGroupName"`
}

func (ListByResourceGroupResource) Uri() string {
	return "/aks/snapshot/listByResourceGroup"
}
func (ListByResourceGroupResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &ListByResourceGroupSnapshot{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	items := SnapshotsClient.NewListByResourceGroupPager(inputRequest.ResourceGroupName, nil)
	result := make([]armcontainerservice.SnapshotsClientListByResourceGroupResponse, 0)
	for items.More() {
		tmp, err := items.NextPage(ctx)
		if err != nil {
			klog.Errorln(err)
		}

		result = append(result, tmp)
	}
	return apis.Response{Code: 200, Data: result}
}
