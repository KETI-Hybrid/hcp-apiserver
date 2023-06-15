package nodepool

import (
	"encoding/json"
	"fmt"
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type CompleteUpgradeResource struct {
	apis.DeleteNotSupported
	apis.GetNotSupported
	apis.PutNotSupported
}

type CompleteUpgrade struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
	NodePoolName string `json:"nodePoolName"`
}

func (CompleteUpgradeResource) Uri() string {
	return "/gke/locations/cluster/nodePools/completeUpgrade"
}

func (CompleteUpgradeResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	gclient := types.GetGKEClient()
	token := gclient.Token
	client := &http.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &CompleteUpgrade{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	// The name (project, location, cluster id) of the cluster to complete IP
	// rotation. Specified in the format 'projects/*/locations/*/clusters/*'.
	url := fmt.Sprintf("https://container.googleapis.com/v1/projects/%s/locations/%s/clusters/%s/nodePools/%s:completeUpgrade", inputRequest.ProjectName, inputRequest.LocationName, inputRequest.ClusterName, inputRequest.NodePoolName) // TODO: Update placeholder value.
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		klog.Errorln(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
