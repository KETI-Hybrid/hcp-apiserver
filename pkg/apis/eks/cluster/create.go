package cluster

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
func (CreateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ClusterCreateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.CreateClusterInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		EncryptionConfig:   make([]*eks.EncryptionConfig, 0),
		KubernetesNetworkConfig: &eks.KubernetesNetworkConfigRequest{
			IpFamily:        aws.String(inputReq.KubernetesNetworkConfig.IpFamily),
			ServiceIpv4Cidr: aws.String(inputReq.KubernetesNetworkConfig.ServiceIpv4Cidr),
		},
		Logging: &eks.Logging{
			ClusterLogging: make([]*eks.LogSetup, 0),
		},
		Name: aws.String(inputReq.Name),
		OutpostConfig: &eks.OutpostConfigRequest{
			ControlPlaneInstanceType: aws.String(inputReq.OutpostConfig.ControlPlaneInstanceType),
			ControlPlanePlacement: &eks.ControlPlanePlacementRequest{
				GroupName: aws.String(inputReq.OutpostConfig.ControlPlanePlacement.GroupName),
			},
			OutpostArns: aws.StringSlice(inputReq.OutpostConfig.OutpostArns),
		},
		ResourcesVpcConfig: &eks.VpcConfigRequest{
			EndpointPrivateAccess: aws.Bool(inputReq.ResourcesVpcConfig.EndpointPrivateAccess),
			EndpointPublicAccess:  aws.Bool(inputReq.ResourcesVpcConfig.EndpointPublicAccess),
			PublicAccessCidrs:     aws.StringSlice(inputReq.ResourcesVpcConfig.PublicAccessCidrs),
			SecurityGroupIds:      aws.StringSlice(inputReq.ResourcesVpcConfig.SecurityGroupIds),
			SubnetIds:             aws.StringSlice(inputReq.ResourcesVpcConfig.SubnetIds),
		},
		RoleArn: aws.String(inputReq.RoleArn),
		Tags:    aws.StringMap(inputReq.Tags),
		Version: aws.String(inputReq.Version),
	}
	for _, encryptionConfig := range inputReq.EncryptionConfig {
		realInput.EncryptionConfig = append(realInput.EncryptionConfig, &eks.EncryptionConfig{
			Provider: &eks.Provider{
				KeyArn: aws.String(encryptionConfig.Provider.KeyArn),
			},
			Resources: aws.StringSlice(encryptionConfig.Resources),
		})
	}
	for _, clusterLogging := range inputReq.Logging.ClusterLogging {
		realInput.Logging.ClusterLogging = append(realInput.Logging.ClusterLogging, &eks.LogSetup{
			Enabled: aws.Bool(clusterLogging.Enabled),
			Types:   aws.StringSlice(clusterLogging.Types),
		})
	}

	result, err := ClusterClient.CreateCluster(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
