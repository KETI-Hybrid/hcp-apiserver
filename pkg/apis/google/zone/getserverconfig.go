package zone

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

type GetServerConfigResource struct {
	apis.PutNotSupported
	apis.DeleteNotSupported
	apis.PostNotSupported
}

type GetServerConfig struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
}

func (GetServerConfigResource) Uri() string {
	return "/gke/zone/getServerconfig"
}

func (GetServerConfigResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetGKEClient()
	containerService := client.ContanerService
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &GetServerConfig{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	// name := fmt.Sprintf("projects/%s/locations/%s", inputRequest.ProjectName, inputRequest.LocationName) // TODO: Update placeholder value.

	resp, err := containerService.Projects.Zones.GetServerconfig(inputRequest.ProjectName, inputRequest.LocationName).Context(ctx).Do()
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
