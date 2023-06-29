package cluster

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

type ClusterCreateInput struct {
	ClientRequestToken      string                         `json:"clientRequestToken"`
	EncryptionConfig        []EncryptionConfig             `json:"encryptionConfig"`
	KubernetesNetworkConfig KubernetesNetworkConfigRequest `json:"kubernetesNetworkConfig"`
	Logging                 Logging                        `json:"logging"`
	Name                    string                         `json:"name"`
	OutpostConfig           OutpostConfigRequest           `json:"outpostConfig"`
	ResourcesVpcConfig      VpcConfigRequest               `json:"resourcesVpcConfig"`
	RoleArn                 string                         `json:"roleArn"`
	Tags                    map[string]string              `json:"tags"`
	Version                 string                         `json:"version"`
}
type EncryptionConfig struct {
	Provider  Provider `json:"provider"`
	Resources []string `json:"resources"`
}
type KubernetesNetworkConfigRequest struct {
	IpFamily        string `json:"ipFamily"`
	ServiceIpv4Cidr string `json:"serviceIpv4Cidr"`
}
type Logging struct {
	ClusterLogging []LogSetup `json:"clusterLogging"`
}
type LogSetup struct {
	Enabled bool     `json:"enabled"`
	Types   []string `json:"types"`
}
type OutpostConfigRequest struct {
	ControlPlaneInstanceType string                       `json:"controlPlaneInstanceType"`
	ControlPlanePlacement    ControlPlanePlacementRequest `json:"controlPlanePlacement"`
	OutpostArns              []string                     `json:"outpostArns"`
}
type ControlPlanePlacementRequest struct {
	GroupName string `json:"groupName"`
}
type VpcConfigRequest struct {
	EndpointPrivateAccess bool     `json:"endpointPrivateAccess"`
	EndpointPublicAccess  bool     `json:"endpointPublicAccess"`
	PublicAccessCidrs     []string `json:"publicAccessCidrs"`
	SecurityGroupIds      []string `json:"securityGroupIds"`
	SubnetIds             []string `json:"subnetIds"`
}
type Provider struct {
	KeyArn string
}

func (CreateResource) Uri() string {
	return "/eks/cluster/create"
}
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(ClusterCreateInput{}, eks.CreateClusterOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
