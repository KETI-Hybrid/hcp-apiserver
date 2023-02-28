package cluster

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
	"k8s.io/klog"
)

type ResourceLabelsResource struct {
	apis.DeleteNotSupported
	apis.GetNotSupported
	apis.PutNotSupported
}

type ResourceLabels struct {
	ProjectName      string            `json:"projectName"`
	LocationName     string            `json:"locationName"`
	ClusterName      string            `json:"clusterName"`
	ResourceLabels   map[string]string `json:"resourceLabels"`
	LabelFingerprint string            `json:"labelFingerprint"`
}

func (ResourceLabelsResource) Uri() string {
	return "/gke/zone/cluster/setResourceLabels"
}

func (ResourceLabelsResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &ResourceLabels{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	//name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName) // TODO: Update placeholder value.
	realRequest := &container.SetLabelsRequest{
		ResourceLabels:   inputRequest.ResourceLabels,
		LabelFingerprint: inputRequest.LabelFingerprint,
	}
	resp, err := containerService.Projects.Zones.Clusters.ResourceLabels(inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, realRequest).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
