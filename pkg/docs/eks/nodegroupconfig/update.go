package nodegroupconfig

import (
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type UpdateResource struct {
	docs.GetNotSupported
	docs.PutNotSupported
	docs.DeleteNotSupported
}

type NodeGroupConfigUpdateInput struct {
	ClientRequestToken string                 `json:"clientRequestToken"`
	ClusterName        string                 `json:"clusterName"`
	Labels             UpdateLabelsPayload    `json:"labels"`
	NodegroupName      string                 `json:"nodegroupName"`
	ScalingConfig      NodegroupScalingConfig `json:"scalingConfig"`
	Taints             UpdateTaintsPayload    `json:"taints"`
	UpdateConfig       NodegroupUpdateConfig  `json:"updateConfig"`
}

type UpdateLabelsPayload struct {
	AddOrUpdateLabels map[string]string `json:"addOrUpdateLabels"`
	RemoveLabels      []string          `json:"removeLabels"`
}

type NodegroupScalingConfig struct {
	DesiredSize int64 `json:"desiredSize"`
	MaxSize     int64 `json:"maxSize"`
	MinSize     int64 `json:"minSize"`
}

type UpdateTaintsPayload struct {
	AddOrUpdateTaints []Taint `json:"addOrUpdateTaints"`
	RemoveTaints      []Taint `json:"removeTaints"`
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

func (UpdateResource) Uri() string {
	return "/eks/nodegroup-config/update"
}
func (UpdateResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &NodeGroupConfigUpdateInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.UpdateNodegroupConfigInput{
		ClientRequestToken: aws.String(inputReq.ClientRequestToken),
		ClusterName:        aws.String(inputReq.ClusterName),
		Labels: &eks.UpdateLabelsPayload{
			AddOrUpdateLabels: aws.StringMap(inputReq.Labels.AddOrUpdateLabels),
			RemoveLabels:      aws.StringSlice(inputReq.Labels.RemoveLabels),
		},
		NodegroupName: aws.String(inputReq.NodegroupName),
		ScalingConfig: &eks.NodegroupScalingConfig{
			DesiredSize: aws.Int64(inputReq.ScalingConfig.DesiredSize),
			MaxSize:     aws.Int64(inputReq.ScalingConfig.MaxSize),
			MinSize:     aws.Int64(inputReq.ScalingConfig.MinSize),
		},
		Taints: &eks.UpdateTaintsPayload{
			AddOrUpdateTaints: make([]*eks.Taint, 0),
			RemoveTaints:      make([]*eks.Taint, 0),
		},
		UpdateConfig: &eks.NodegroupUpdateConfig{
			MaxUnavailable:           aws.Int64(inputReq.UpdateConfig.MaxUnavailable),
			MaxUnavailablePercentage: aws.Int64(inputReq.UpdateConfig.MaxUnavailablePercentage),
		},
	}

	for _, taint := range inputReq.Taints.AddOrUpdateTaints {
		realInput.Taints.AddOrUpdateTaints = append(realInput.Taints.AddOrUpdateTaints, &eks.Taint{
			Effect: aws.String(taint.Effect),
			Key:    aws.String(taint.Key),
			Value:  aws.String(taint.Value),
		})
	}

	for _, taint := range inputReq.Taints.RemoveTaints {
		realInput.Taints.RemoveTaints = append(realInput.Taints.RemoveTaints, &eks.Taint{
			Effect: aws.String(taint.Effect),
			Key:    aws.String(taint.Key),
			Value:  aws.String(taint.Value),
		})
	}

	result, err := NodeGroupConfigClient.UpdateNodegroupConfig(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return docs.Response{Code: 200, Data: result}
}
