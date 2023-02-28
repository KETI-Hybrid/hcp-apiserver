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

type AutoScalingResource struct {
	apis.DeleteNotSupported
	apis.GetNotSupported
	apis.PutNotSupported
}

type AutoScaling struct {
	ProjectName  string                         `json:"projectName"`
	LocationName string                         `json:"locationName"`
	ClusterName  string                         `json:"clusterName"`
	NodePoolName string                         `json:"nodePoolName"`
	NodePool     *container.NodePoolAutoscaling `json:"nodePool"`
}

func (AutoScalingResource) Uri() string {
	return "/gke/zone/cluster/nodePools/autoscaling"
}

func (AutoScalingResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &AutoScaling{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName) // TODO: Update placeholder value.
	realRequest := &container.SetNodePoolAutoscalingRequest{
		Name:        name,
		Autoscaling: inputRequest.NodePool,
	}
	resp, err := containerService.Projects.Zones.Clusters.NodePools.Autoscaling(
		inputRequest.ProjectName,
		inputRequest.LocationName,
		inputRequest.ClusterName,
		inputRequest.NodePoolName,
		realRequest).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
