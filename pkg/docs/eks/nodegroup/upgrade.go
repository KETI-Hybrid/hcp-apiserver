package nodegroup

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

type NodeGroupUpgradeInput struct {
	ClientRequestToken string                      `json:"clientRequestToken"`
	ClusterName        string                      `json:"clusterName"`
	Force              bool                        `json:"force"`
	LaunchTemplate     LaunchTemplateSpecification `json:"launchTemplate"`
	NodegroupName      string                      `json:"nodegroupName"`
	ReleaseVersion     string                      `json:"releaseVersion"`
	Version            string                      `json:"version"`
}

func (UpgradeResource) Uri() string {
	return "/eks/nodegroup/upgrade"
}
func (UpgradeResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(NodeGroupUpgradeInput{}, eks.UpdateNodegroupVersionOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
