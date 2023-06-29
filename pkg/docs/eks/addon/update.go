package addon

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type UpdateResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type AddonUpdateInput struct {
	AddonName             string `json:"addonName"`
	AddonVersion          string `json:"addonVersion"`
	ClientRequestToken    string `json:"clientRequestToken"`
	ClusterName           string `json:"clusterName"`
	ConfigurationValues   string `json:"configurationValues"`
	ResolveConflicts      string `json:"resolveConflicts"`
	ServiceAccountRoleArn string `json:"serviceAccountRoleArn"`
}

func (UpdateResource) Uri() string {
	return "/eks/addon/update"
}
func (UpdateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(AddonUpdateInput{}, eks.UpdateAddonOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
