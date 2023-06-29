package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
)

type RegisterResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}
type ClusterRegisterInput struct {
	ClientRequestToken string                 `json:"clientRequestToken"`
	ConnectorConfig    ConnectorConfigRequest `json:"connectorConfig"`
	Name               string                 `json:"name"`
	Tags               map[string]string      `json:"tags"`
}
type ConnectorConfigRequest struct {
	Provider string `json:"provider"`
	RoleArn  string `json:"roleArn"`
}

func (RegisterResource) Uri() string {
	return "/eks/cluster/register"
}
func (RegisterResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(ClusterRegisterInput{}, eks.RegisterClusterOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
