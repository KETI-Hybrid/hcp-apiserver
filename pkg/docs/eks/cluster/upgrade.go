package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type UpgradeResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type ClusterUpgradeInput struct {
	ClientRequestToken string `json:"clientRequestToken"`
	Name               string `json:"name"`
	Version            string `json:"version"`
}

func (UpgradeResource) Uri() string {
	return "/eks/cluster/upgrade"
}
func (UpgradeResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(ClusterUpgradeInput{}, eks.UpdateClusterVersionOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
