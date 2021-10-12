package handler

import "github.com/aws/aws-sdk-go/service/eks"

func ListUpdate(listUpdateInput eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {

	eksSvc := GetEKSClient(listUpdateInput.Name)

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

	eksSvc := GetEKSClient(describeUpdateInput.Name)

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

	eksSvc := GetEKSClient(input.Name)

	input = eks.UpdateClusterConfigInput{
		ClientRequestToken: input.ClientRequestToken,
		Logging:            input.Logging,
		Name:               input.Name,
		ResourcesVpcConfig: input.ResourcesVpcConfig,
	}
	out, err := eksSvc.UpdateClusterConfig(&input)

	return out, err
}

func UpdateNodegroupConfig(input eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {

	eksSvc := GetEKSClient(input.ClusterName)

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
