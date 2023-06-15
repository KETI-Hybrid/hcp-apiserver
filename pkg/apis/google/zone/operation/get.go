package operation

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type GetResource struct {
	apis.PutNotSupported
	apis.DeleteNotSupported
	apis.PostNotSupported
}

type Get struct {
	ProjectName   string `json:"projectName"`
	LocationName  string `json:"locationName"`
	Operationname string `json:"operationName"`
}

func (GetResource) Uri() string {
	return "/gke/zone/operation/get"
}

func (GetResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Cancel{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	// name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.Operationname) // TODO: Update placeholder value.

	resp, err := containerService.Projects.Zones.Operations.Get(inputRequest.ProjectName, inputRequest.LocationName, inputRequest.Operationname).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
