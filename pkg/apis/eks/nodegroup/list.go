package nodegroup

import (
	"hcp-apiserver/pkg/apis"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ListResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}

type NodeGroupListInput struct {
	ClusterName string `json:"clusterName"`
	MaxResults  int64  `json:"maxResults"`
	NextToken   string `json:"nextToken"`
}

func (ListResource) Uri() string {
	return "/eks/nodegroup/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	return apis.Response{Code: 200, Data: nil}
}
