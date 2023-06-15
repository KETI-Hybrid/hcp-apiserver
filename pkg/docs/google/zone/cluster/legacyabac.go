package cluster

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
	"k8s.io/klog"
)

type LegacyAbacResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type LegacyAbac struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
	Enabled      bool   `json:"enabled"`
}

func (LegacyAbacResource) Uri() string {
	return "/gke/zone/cluster/legacyAbac"
}

func (LegacyAbacResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &LegacyAbac{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	//name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName) // TODO: Update placeholder value.
	realRequest := &container.SetLegacyAbacRequest{
		Enabled: inputRequest.Enabled,
	}
	resp, err := containerService.Projects.Zones.Clusters.LegacyAbac(inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, realRequest).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: resp}
}
