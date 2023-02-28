package nodepool

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

type SetManagementResource struct {
	apis.GetNotSupported
	apis.DeleteNotSupported
	apis.PutNotSupported
}

type Management struct {
	ProjectName  string                    `json:"projectName"`
	LocationName string                    `json:"locationName"`
	ClusterName  string                    `json:"clusterName"`
	NodePoolName string                    `json:"nodepoolName"`
	Management   *container.NodeManagement `json:"management"`
}

func (SetManagementResource) Uri() string {
	return "/gke/zone/cluster/nodePools/setManagement"
}

func (SetManagementResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Management{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	//name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, inputRequest.NodePoolName) // TODO: Update placeholder value.
	realRequest := &container.SetNodePoolManagementRequest{
		Management: inputRequest.Management,
	}
	resp, err := containerService.Projects.Zones.Clusters.NodePools.SetManagement(inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, inputRequest.NodePoolName, realRequest).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
