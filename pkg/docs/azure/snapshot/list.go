package snapshot

import (
	"context"
	"hcp-apiserver/pkg/docs"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListResource struct {
	docs.DeleteNotSupported
	docs.PostNotSupported
	docs.PutNotSupported
}

func (ListResource) Uri() string {
	return "/aks/snapshot/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	ctx := context.Background()
	items := SnapshotsClient.NewListPager(nil)
	result := make([]armcontainerservice.SnapshotsClientListResponse, 0)
	for items.More() {
		tmp, err := items.NextPage(ctx)
		if err != nil {
			klog.Errorln(err)
		}

		result = append(result, tmp)
	}
	return docs.Response{Code: 200, Data: result}
}
