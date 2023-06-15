package usablesubnet

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

type ListResource struct {
	apis.DeleteNotSupported
	apis.PostNotSupported
	apis.PutNotSupported
}

type List struct {
	ProjectName string `json:"projectName"`
}

func (ListResource) Uri() string {
	return "/gke/aggregated/usableSubnetworks/list"
}

func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
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
	name := fmt.Sprintf("projects/%s", inputRequest.ProjectName) // TODO: Update placeholder value.
	resp, err := containerService.Projects.Aggregated.UsableSubnetworks.List(name).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
