package privatelink

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type ListPrivateLink struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
}

func (ListResource) Uri() string {
	return "/aks/privateLink/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &ListPrivateLink{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	result, err := PrivateLinkClient.List(ctx, inputRequest.ResourceGroupName, inputRequest.ClusterName, nil)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
