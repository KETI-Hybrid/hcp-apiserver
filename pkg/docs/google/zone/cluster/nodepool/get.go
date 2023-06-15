package nodepool

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type GetResource struct {
	docs.PostNotSupported
	docs.DeleteNotSupported
	docs.PutNotSupported
}

type Get struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
	NodePoolName string `json:"nodepoolName"`
}

func (GetResource) Uri() string {
	return "/gke/zone/cluster/nodePools/get"
}

func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
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
	//name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, inputRequest.NodePoolName) // TODO: Update placeholder value.

	resp, err := containerService.Projects.Zones.Clusters.NodePools.Get(inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, inputRequest.NodePoolName).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: resp}
}
