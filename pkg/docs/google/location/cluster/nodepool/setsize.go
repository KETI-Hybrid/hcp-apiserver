package nodepool

import (
	"context"
	"encoding/json"
	"fmt"
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
	"k8s.io/klog"
)

type SetSizeResource struct {
	apis.GetNotSupported
	apis.DeleteNotSupported
	apis.PutNotSupported
}

type Size struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
	NodePoolName string `json:"nodepoolName"`
	NodeCount    int64  `json:"nodeCount"`
}

func (SetSizeResource) Uri() string {
	return "/gke/locations/cluster/nodePools/setSize"
}

func (SetSizeResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Size{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, inputRequest.NodePoolName) // TODO: Update placeholder value.
	realRequest := &container.SetNodePoolSizeRequest{
		NodeCount: inputRequest.NodeCount,
	}
	resp, err := containerService.Projects.Locations.Clusters.NodePools.SetSize(name, realRequest).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
