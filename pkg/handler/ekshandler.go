package handler

import (
	cobrautil "Hybrid_Cluster/hybridctl/util"
	clusterRegister "Hybrid_Cluster/pkg/client/clusterregister/v1alpha1/clientset/versioned/typed/clusterregister/v1alpha1"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// addon

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetEKSClient(clusterName *string) (*eks.EKS, error) {
	master_config, _ := cobrautil.BuildConfigFromFlags("kube-master", "/root/.kube/config")
	clusterRegisterClientSet, err := clusterRegister.NewForConfig(master_config)
	checkErr(err)
	clusterRegisters, err := clusterRegisterClientSet.ClusterRegisters("eks").Get(context.TODO(), *clusterName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(clusterRegisters.Spec.Region),
	}))
	eksSvc := eks.New(sess)
	return eksSvc, nil
}

func CreateAddon(addonInput eks.CreateAddonInput) (*eks.CreateAddonOutput, error) {

	// println(*addonInput.ClusterName)
	eksSvc, err := GetEKSClient(addonInput.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newAddonInput := &eks.CreateAddonInput{
		AddonName:             addonInput.AddonName,
		AddonVersion:          addonInput.AddonVersion,
		ClientRequestToken:    addonInput.ClientRequestToken,
		ClusterName:           addonInput.ClusterName,
		ResolveConflicts:      addonInput.ResolveConflicts,
		ServiceAccountRoleArn: addonInput.ServiceAccountRoleArn,
		Tags:                  addonInput.Tags,
	}
	out, err := eksSvc.CreateAddon(newAddonInput)

	return out, err
}

func DeleteAddon(addonInput eks.DeleteAddonInput) (*eks.DeleteAddonOutput, error) {

	eksSvc, err := GetEKSClient(addonInput.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newAddonInput := &eks.DeleteAddonInput{
		AddonName:   addonInput.AddonName,
		ClusterName: addonInput.ClusterName,
	}
	out, err := eksSvc.DeleteAddon(newAddonInput)

	return out, err
}

func DescribeAddon(addonInput eks.DescribeAddonInput) (*eks.DescribeAddonOutput, error) {

	eksSvc, err := GetEKSClient(addonInput.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newAddonInput := &eks.DescribeAddonInput{
		AddonName:   addonInput.AddonName,
		ClusterName: addonInput.ClusterName,
	}
	out, err := eksSvc.DescribeAddon(newAddonInput)

	return out, err
}

func DescribeAddonVersions(addonInput eks.DescribeAddonVersionsInput) (*eks.DescribeAddonVersionsOutput, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	newAddonInput := &eks.DescribeAddonVersionsInput{
		AddonName: addonInput.AddonName,
	}
	out, err := eksSvc.DescribeAddonVersions(newAddonInput)

	return out, err
}

func ListAddon(addonInput eks.ListAddonsInput) (*eks.ListAddonsOutput, error) {

	eksSvc, err := GetEKSClient(addonInput.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newAddonInput := &eks.ListAddonsInput{
		ClusterName: addonInput.ClusterName,
	}
	out, err := eksSvc.ListAddons(newAddonInput)

	return out, err
}

func UpdateAddon(addonInput eks.UpdateAddonInput) (*eks.UpdateAddonOutput, error) {

	eksSvc, err := GetEKSClient(addonInput.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newAddonInput := &eks.UpdateAddonInput{
		ClusterName: addonInput.ClusterName,
		AddonName:   addonInput.AddonName,
	}
	out, err := eksSvc.UpdateAddon(newAddonInput)

	return out, err
}

func AssociateEncryptionConfig(input eks.AssociateEncryptionConfigInput) (*eks.AssociateEncryptionConfigOutput, error) {
	eksSvc, err := GetEKSClient(input.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newInput := &eks.AssociateEncryptionConfigInput{
		ClientRequestToken: input.ClientRequestToken,
		ClusterName:        input.ClusterName,
		EncryptionConfig:   input.EncryptionConfig,
	}
	out, err := eksSvc.AssociateEncryptionConfig(newInput)

	return out, err
}

// identity provider

func AssociateIdentityProviderConfig(input eks.AssociateIdentityProviderConfigInput) (*eks.AssociateIdentityProviderConfigOutput, error) {

	// println(*Input.ClusterName)
	eksSvc, err := GetEKSClient(input.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newInput := &eks.AssociateIdentityProviderConfigInput{
		ClientRequestToken: input.ClientRequestToken,
		ClusterName:        input.ClusterName,
		Oidc:               input.Oidc,
		Tags:               input.Tags,
	}
	out, err := eksSvc.AssociateIdentityProviderConfig(newInput)

	return out, err
}

func DisassociateIdentityProviderConfig(input eks.DisassociateIdentityProviderConfigInput) (*eks.DisassociateIdentityProviderConfigOutput, error) {

	// println(*Input.ClusterName)
	eksSvc, err := GetEKSClient(input.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newInput := &eks.DisassociateIdentityProviderConfigInput{
		ClientRequestToken:     input.ClientRequestToken,
		ClusterName:            input.ClusterName,
		IdentityProviderConfig: input.IdentityProviderConfig,
	}
	out, err := eksSvc.DisassociateIdentityProviderConfig(newInput)

	return out, err
}

func DescribeIdentityProviderConfig(input eks.DescribeIdentityProviderConfigInput) (*eks.DescribeIdentityProviderConfigOutput, error) {

	// println(*Input.ClusterName)
	eksSvc, err := GetEKSClient(input.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newInput := &eks.DescribeIdentityProviderConfigInput{
		ClusterName:            input.ClusterName,
		IdentityProviderConfig: input.IdentityProviderConfig,
	}
	out, err := eksSvc.DescribeIdentityProviderConfig(newInput)

	return out, err
}

func ListIdentityProviderConfigs(input eks.ListIdentityProviderConfigsInput) (*eks.ListIdentityProviderConfigsOutput, error) {

	// println(*Input.ClusterName)
	eksSvc, err := GetEKSClient(input.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newInput := &eks.ListIdentityProviderConfigsInput{
		ClusterName: input.ClusterName,
		MaxResults:  input.MaxResults,
		NextToken:   input.NextToken,
	}
	out, err := eksSvc.ListIdentityProviderConfigs(newInput)

	return out, err
}

// tag

func ListTagsForResource(listTagsForResourceInput eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	input := &eks.ListTagsForResourceInput{
		ResourceArn: listTagsForResourceInput.ResourceArn,
	}
	out, err := eksSvc.ListTagsForResource(input)

	return out, err
}

func TagResource(input eks.TagResourceInput) (*eks.TagResourceOutput, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	input = eks.TagResourceInput{
		ResourceArn: input.ResourceArn,
		Tags:        input.Tags,
	}
	out, err := eksSvc.TagResource(&input)

	return out, err
}

func UntagResource(input eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	fmt.Println(input.TagKeys)
	input = eks.UntagResourceInput{
		ResourceArn: input.ResourceArn,
		TagKeys:     input.TagKeys,
	}
	out, err := eksSvc.UntagResource(&input)

	return out, err
}

// update

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
