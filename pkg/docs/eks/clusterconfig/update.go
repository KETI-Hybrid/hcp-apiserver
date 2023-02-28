package clusterconfig

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

type ClusterConfigUpdateInput struct {
	ClientRequestToken string           `json:"clientRequestToken"`
	Logging            Logging          `json:"logging"`
	Name               string           `json:"name"`
	ResourcesVpcConfig VpcConfigRequest `json:"resourcesVpcConfig"`
}
type Logging struct {
	ClusterLogging []LogSetup `json:"clusterLogging"`
}
type LogSetup struct {
	Enabled bool     `json:"enabled"`
	Types   []string `json:"types"`
}
type VpcConfigRequest struct {
	EndpointPrivateAccess bool     `json:"endpointPrivateAccess"`
	EndpointPublicAccess  bool     `json:"endpointPublicAccess"`
	PublicAccessCidrs     []string `json:"publicAccessCidrs"`
	SecurityGroupIds      []string `json:"securityGroupIds"`
	SubnetIds             []string `json:"subnetIds"`
}

func (UpdateResource) Uri() string {
	return "/eks/cluster-config/update"
}
func (UpdateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(ClusterConfigUpdateInput{}, eks.UpdateClusterConfigOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
