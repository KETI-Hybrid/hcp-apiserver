package managedcluster

import (
	"context"
	"hcp-apiserver/pkg/apis"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListResource struct {
	apis.DeleteNotSupported
	apis.PostNotSupported
	apis.PutNotSupported
}

func (ListResource) Uri() string {
	return "/aks/managedClusters/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	ctx := context.Background()
	items := ManagedClustersClient.NewListPager(nil)
	result := make([]armcontainerservice.ManagedClustersClientListResponse, 0)
	for items.More() {
		tmp, err := items.NextPage(ctx)
		if err != nil {
			klog.Errorln(err)
		}

		result = append(result, tmp)
	}
	return apis.Response{Code: 200, Data: result}
}
