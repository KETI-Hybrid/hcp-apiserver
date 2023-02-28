package update

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type ListResource struct {
	docs.PostNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type UpdateListInput struct {
	AddonName     string `json:"addonName"`
	MaxResults    int64  `json:"maxResults"`
	Name          string `json:"name"`
	NextToken     string `json:"nextToken"`
	NodegroupName string `json:"nodegroupName"`
}

func (ListResource) Uri() string {
	return "/eks/update/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(UpdateListInput{}, eks.ListUpdatesOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
