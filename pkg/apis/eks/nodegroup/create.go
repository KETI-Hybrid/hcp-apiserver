package nodegroup

import (
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type CreateResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
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
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &NodeGroupCreateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.CreateNodegroupInput{
		AmiType:            aws.String(inputReq.AmiType),
		CapacityType:       aws.String(inputReq.CapacityType),
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		ClusterName:        aws.String(inputReq.ClusterName),
		DiskSize:           aws.Int64(inputReq.DiskSize),
		InstanceTypes:      aws.StringSlice(inputReq.InstanceTypes),
		Labels:             aws.StringMap(inputReq.Labels),
		LaunchTemplate: &eks.LaunchTemplateSpecification{
			Id:      aws.String(inputReq.LaunchTemplate.Id),
			Name:    aws.String(inputReq.LaunchTemplate.Name),
			Version: aws.String(inputReq.LaunchTemplate.Version),
		},
		NodeRole:       aws.String(inputReq.NodeRole),
		NodegroupName:  aws.String(inputReq.NodegroupName),
		ReleaseVersion: aws.String(inputReq.ReleaseVersion),
		RemoteAccess: &eks.RemoteAccessConfig{
			Ec2SshKey:            aws.String(inputReq.RemoteAccess.Ec2SshKey),
			SourceSecurityGroups: aws.StringSlice(inputReq.RemoteAccess.SourceSecurityGroups),
		},
		ScalingConfig: &eks.NodegroupScalingConfig{
			DesiredSize: aws.Int64(inputReq.ScalingConfig.DesiredSize),
			MaxSize:     aws.Int64(inputReq.ScalingConfig.MaxSize),
			MinSize:     aws.Int64(inputReq.ScalingConfig.MinSize),
		},
		Subnets: aws.StringSlice(inputReq.Subnets),
		Tags:    aws.StringMap(inputReq.Tags),
		Taints:  make([]*eks.Taint, 0),
		UpdateConfig: &eks.NodegroupUpdateConfig{
			MaxUnavailable:           aws.Int64(inputReq.UpdateConfig.MaxUnavailable),
			MaxUnavailablePercentage: aws.Int64(inputReq.UpdateConfig.MaxUnavailablePercentage),
		},
	}
	for _, taint := range inputReq.Taints {
		realInput.Taints = append(realInput.Taints, &eks.Taint{
			Effect: aws.String(taint.Effect),
			Key:    aws.String(taint.Key),
			Value:  aws.String(taint.Value),
		})
	}

	result, err := NodeGroupClient.CreateNodegroup(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
