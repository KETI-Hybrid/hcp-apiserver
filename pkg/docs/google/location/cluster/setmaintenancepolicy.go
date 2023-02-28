package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
	"k8s.io/klog"
)

type SetMaintenancePolicyResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type SetMaintenancePolicy struct {
	ProjectName       string                       `json:"projectName"`
	LocationName      string                       `json:"locationName"`
	ClusterName       string                       `json:"clusterName"`
	MaintenancePolicy *container.MaintenancePolicy `json:"maintenancePolicy"`
}

func (SetMaintenancePolicyResource) Uri() string {
	return "/gke/locations/cluster/setMaintenancePolicy"
}

func (SetMaintenancePolicyResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &SetMaintenancePolicy{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName) // TODO: Update placeholder value.
	realRequest := &container.SetMaintenancePolicyRequest{
		MaintenancePolicy: inputRequest.MaintenancePolicy,
	}
	resp, err := containerService.Projects.Locations.Clusters.SetMaintenancePolicy(name, realRequest).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: resp}
}
