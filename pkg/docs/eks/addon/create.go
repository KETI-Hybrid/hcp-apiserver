package addon

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type CreateResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type AddonCreateInput struct {
	AddonName             string            `json:"addonName"`
	AddonVersion          string            `json:"addonVersion"`
	ClientRequestToken    string            `json:"clientRequestToken"`
	ClusterName           string            `json:"clusterName"`
	ConfigurationValues   string            `json:"configurationValues"`
	ResolveConflicts      string            `json:"resolveConflicts"`
	ServiceAccountRoleArn string            `json:"serviceAccountRoleArn"`
	Tags                  map[string]string `json:"tags"`
}

func (CreateResource) Uri() string {
	return "/eks/addon/create"
}
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(AddonCreateInput{}, eks.CreateAddonOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
