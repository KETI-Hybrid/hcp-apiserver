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
	"k8s.io/klog"
)

type GetResource struct {
	apis.PostNotSupported
	apis.DeleteNotSupported
	apis.PutNotSupported
}

type Get struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
}

func (GetResource) Uri() string {
	return "/gke/locations/cluster/get"
}

func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Get{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName) // TODO: Update placeholder value.

	resp, err := containerService.Projects.Locations.Clusters.Get(name).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
