package snapshot

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type UpdateTagsResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.GetNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type UpdateTagSnapshot struct {
	ResourceGroupName string                          `json:"resourceGroupName"`
	ClusterName       string                          `json:"clusterName"`
	Tags              *armcontainerservice.TagsObject `json:"tags"`
}

func (UpdateTagsResource) Uri() string {
	return "/aks/snapshot/updateTag"
}
func (UpdateTagsResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &UpdateTagSnapshot{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := SnapshotsClient.UpdateTags(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, *inputRequest.Tags, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
