package clusterconfig

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

type UpdateResource struct {
	apis.GetNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
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
func (UpdateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &ClusterConfigUpdateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.UpdateClusterConfigInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		Logging: &eks.Logging{
			ClusterLogging: make([]*eks.LogSetup, 0),
		},
		Name: aws.String(inputReq.Name),
		ResourcesVpcConfig: &eks.VpcConfigRequest{
			EndpointPrivateAccess: aws.Bool(inputReq.ResourcesVpcConfig.EndpointPrivateAccess),
			EndpointPublicAccess:  aws.Bool(inputReq.ResourcesVpcConfig.EndpointPublicAccess),
			PublicAccessCidrs:     aws.StringSlice(inputReq.ResourcesVpcConfig.PublicAccessCidrs),
			SecurityGroupIds:      aws.StringSlice(inputReq.ResourcesVpcConfig.SecurityGroupIds),
			SubnetIds:             aws.StringSlice(inputReq.ResourcesVpcConfig.SubnetIds),
		},
	}
	for _, clusterLogging := range inputReq.Logging.ClusterLogging {
		realInput.Logging.ClusterLogging = append(realInput.Logging.ClusterLogging, &eks.LogSetup{
			Enabled: aws.Bool(clusterLogging.Enabled),
			Types:   aws.StringSlice(clusterLogging.Types),
		})
	}

	result, err := ClusterConfigClient.UpdateClusterConfig(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
