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
	"k8s.io/klog"
)

type ListResource struct {
	docs.PostNotSupported
	docs.DeleteNotSupported
	docs.PutNotSupported
}

type List struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
}

func (ListResource) Uri() string {
	return "/gke/locations/cluster/list"
}

func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &List{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	name := fmt.Sprintf("projects/%s/locations/%s", inputRequest.ProjectName, inputRequest.LocationName) // TODO: Update placeholder value.

	resp, err := containerService.Projects.Locations.Clusters.List(name).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: resp}
}
