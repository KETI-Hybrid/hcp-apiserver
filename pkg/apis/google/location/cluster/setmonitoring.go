package cluster

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

type SetMonitoringResource struct {
	apis.DeleteNotSupported
	apis.GetNotSupported
	apis.PutNotSupported
}

type SetMonitoring struct {
	ProjectName       string `json:"projectName"`
	LocationName      string `json:"locationName"`
	ClusterName       string `json:"clusterName"`
	MonitoringService string `json:"monitoringService"`
}

func (SetMonitoringResource) Uri() string {
	return "/gke/locations/cluster/setMonitoring"
}

func (SetMonitoringResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &SetMonitoring{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName) // TODO: Update placeholder value.
	realRequest := &container.SetMonitoringServiceRequest{
		MonitoringService: inputRequest.MonitoringService,
	}
	resp, err := containerService.Projects.Locations.Clusters.SetMonitoring(name, realRequest).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
