package nodegroupconfig

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
	request, response := util.DocWithReq(NodeGroupConfigUpdateInput{}, eks.UpdateNodegroupConfigOutput{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
