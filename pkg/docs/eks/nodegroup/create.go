package nodegroup

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

type NodeGroupCreateInput struct {
	AmiType            string                      `json:"amiType"`
	CapacityType       string                      `json:"capacityType"`
	ClientRequestToken string                      `json:"clientRequestToken"`
	ClusterName        string                      `json:"clusterName"`
	DiskSize           int64                       `json:"diskSize"`
	InstanceTypes      []string                    `json:"instanceTypes"`
	Labels             map[string]string           `json:"labels"`
	LaunchTemplate     LaunchTemplateSpecification `json:"launchTemplate"`
	NodeRole           string                      `json:"nodeRole"`
	NodegroupName      string                      `json:"nodegroupName"`
	ReleaseVersion     string                      `json:"releaseVersion"`
	RemoteAccess       RemoteAccessConfig          `json:"remoteAccess"`
	ScalingConfig      NodegroupScalingConfig      `json:"scalingConfig"`
	Subnets            []string                    `json:"subnets"`
	Tags               map[string]string           `json:"tags"`
	Taints             []Taint                     `json:"taints"`
	UpdateConfig       NodegroupUpdateConfig       `json:"updateConfig"`
	Version            string                      `json:"version"`
}

type LaunchTemplateSpecification struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}
type RemoteAccessConfig struct {
	Ec2SshKey            string   `json:"ec2SshKey"`
	SourceSecurityGroups []string `json:"sourceSecurityGroups"`
}
type NodegroupScalingConfig struct {
	DesiredSize int64 `json:"desiredSize"`
	MaxSize     int64 `json:"maxSize"`
	MinSize     int64 `json:"minSize"`
}
type Taint struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}
type NodegroupUpdateConfig struct {
	MaxUnavailable           int64 `json:"maxUnavailable"`
	MaxUnavailablePercentage int64 `json:"maxUnavailablePercentage"`
}

func (CreateResource) Uri() string {
	return "/eks/nodegroup/create"
}
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(NodeGroupCreateInput{}, eks.CreateNodegroupOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
