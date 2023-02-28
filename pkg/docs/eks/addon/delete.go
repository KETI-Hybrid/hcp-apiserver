package addon

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type DeleteResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.PostNotSupported
}

type AddonDeleteInput struct {
	AddonName   string `json:"addonName"`
	ClusterName string `json:"clusterName"`
	Preserve    bool   `json:"preserve"`
}

func (DeleteResource) Uri() string {
	return "/eks/addon/delete"
}
func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(AddonDeleteInput{}, eks.DeleteAddonOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
