package cluster

import (
	"context"
	"hcp-apiserver/pkg/apis"
	"hcp-apiserver/pkg/types"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListResource struct {
	apis.PutNotSupported
	apis.DeleteNotSupported
	apis.PostNotSupported
}

func (ListResource) Uri() string {
	return "/nks/cluster/list"
}

func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	client := types.GetNKSClient()
	containerService := client.Client.V2Api
	ctx := context.Background()

	resp, err := containerService.ClustersGet(ctx)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: resp}
}
