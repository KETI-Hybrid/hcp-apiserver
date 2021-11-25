package handler

import "github.com/aws/aws-sdk-go/service/eks"

func ListUpdate(listUpdateInput eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {

	eksSvc, err := GetEKSClient(listUpdateInput.Name)
	if eksSvc == nil {
		return nil, err
	}
	input := &eks.ListUpdatesInput{
		AddonName:     listUpdateInput.AddonName,
		MaxResults:    listUpdateInput.MaxResults,
		Name:          listUpdateInput.Name,
		NextToken:     listUpdateInput.NextToken,
		NodegroupName: listUpdateInput.NodegroupName,
	}
	out, err := eksSvc.ListUpdates(input)

	return out, err
}

func DescribeUpdate(describeUpdateInput eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {

	eksSvc, err := GetEKSClient(describeUpdateInput.Name)
	if eksSvc == nil {
		return nil, err
	}
	input := &eks.DescribeUpdateInput{
		AddonName:     describeUpdateInput.AddonName,
		Name:          describeUpdateInput.Name,
		NodegroupName: describeUpdateInput.NodegroupName,
		UpdateId:      describeUpdateInput.UpdateId,
	}
	out, err := eksSvc.DescribeUpdate(input)

	return out, err
}

func UpdateClusterConfig(input eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error) {

	eksSvc, err := GetEKSClient(input.Name)
	if eksSvc == nil {
		return nil, err
	}
	input = eks.UpdateClusterConfigInput{
		ClientRequestToken: input.ClientRequestToken,
		Logging:            input.Logging,
		Name:               input.Name,
		ResourcesVpcConfig: input.ResourcesVpcConfig,
	}
	out, err := eksSvc.UpdateClusterConfig(&input)

	return out, err
}

func UpdateNodeGroupConfig(input eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {

	eksSvc, err := GetEKSClient(input.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	input = eks.UpdateNodegroupConfigInput{
		ClientRequestToken: input.ClientRequestToken,
		ClusterName:        input.ClusterName,
		Labels:             input.Labels,
		NodegroupName:      input.NodegroupName,
		ScalingConfig:      input.ScalingConfig,
		Taints:             input.Taints,
		UpdateConfig:       input.UpdateConfig,
	}
	out, err := eksSvc.UpdateNodegroupConfig(&input)

	return out, err
}
