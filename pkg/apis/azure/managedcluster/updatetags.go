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

type UpdateTagsResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type UpdateTags struct {
	ResourceGroupName string                          `json:"resourceGroupName"`
	ClusterName       string                          `json:"clusterName"`
	Tags              *armcontainerservice.TagsObject `json:"tags"`
}

func (UpdateTagsResource) Uri() string {
	return "/aks/managedClusters/updateTags"
}
func (UpdateTagsResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &UpdateTags{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := inputRequest.Tags
	ctx := context.Background()
	result, err := ManagedClustersClient.BeginUpdateTags(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, *realInput, nil)
	if err != nil {
		klog.Errorln(err)
	}
	resp, err := result.PollUntilDone(ctx, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
