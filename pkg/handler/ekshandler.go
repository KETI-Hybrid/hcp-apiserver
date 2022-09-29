package handler

import (
	"context"
	"fmt"

	"hybridctl/util"

	"hcp-pkg/util/clusterManager"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

// addon
var cm, _ = clusterManager.NewClusterManager()

func GetEKSClient(clusterName *string) (*eks.EKS, error) {
	cluster, err := cm.HCPCluster_Client.HcpV1alpha1().HCPClusters("hcp").Get(context.TODO(), *clusterName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	klog.Info("[1] ", cluster.Spec.Region)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(cluster.Spec.Region),
	}))

	eksSvc := eks.New(sess)
	return eksSvc, nil
}

func InitializeEKSClient(region string) *eks.EKS {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Credentials: credentials.NewSharedCredentials("/root/.aws/credentials", "default"),
			Region:      aws.String(region),
		},
	}))
	eksSvc := eks.New(sess, aws.NewConfig().WithRegion(region))
	return eksSvc
}

func EKSCreateCluster(Input util.HCPCreateClusterInput) (*eks.CreateClusterOutput, error) {
	klog.Info("Called EKSCreateCluster")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}

	fmt.Println(Input.EKSCreateClusterInput)

	out, err := eksSvc.CreateCluster(&Input.EKSCreateClusterInput)
	fmt.Println(out)
	fmt.Println(err)
	return out, err
}

func EKSDeleteCluster(Input util.HCPDeleteClusterInput) (*eks.DeleteClusterOutput, error) {
	klog.Info("Called EKSDeleteCluster")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}

	out, err := eksSvc.DeleteCluster(&Input.EKSDeleteClusterInput)
	return out, err
}

func EKSDescribeCluster(Input util.HCPDescribeClusterInput) (*eks.DescribeClusterOutput, error) {
	klog.Infoln("Called EksDescribeCluster")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}

	fmt.Println(Input.EKSDescribeClusterInput)

	out, err := eksSvc.DescribeCluster(&Input.EKSDescribeClusterInput)
	return out, err
}

func EKSListCluster(Input util.HCPListClusterInput) (*eks.ListClustersOutput, error) {
	klog.Infoln("Called EKSListCluster")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}

	input := &eks.ListClustersInput{}
	result, err := eksSvc.ListClusters(input)
	if err != nil {
		klog.Infoln(err)
	}
	klog.Infoln(result)

	// out, err := eksSvc.ListClusters(&Input.EKSListClusterInput)
	return result, err
}

func EKSUpgradeCluster(Input util.HCPUpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error) {
	klog.Infoln("Called EKS Cluster Version Upgrade")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}

	fmt.Println(Input.EKSUpdateClusterVersionInput)
	out, err := eksSvc.UpdateClusterVersion(&Input.EKSUpdateClusterVersionInput)
	return out, err
}

func EKSCreateNodegroup(Input util.HCPCreateNodegroupInput) (*eks.CreateNodegroupOutput, error) {
	klog.Info("Called EKSCreateNodegroup")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}

	out, err := eksSvc.CreateNodegroup(&Input.EKSCreateNodegroupInput)
	return out, err
}

func EKSDeleteNodegroup(Input util.HCPDeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error) {
	klog.Info("Called EKSCreateNodegroup")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}

	out, err := eksSvc.DeleteNodegroup(&Input.EKSDeleteNodegroupInput)
	return out, err
}

func EKSDescribeNodegroup(Input util.HCPDescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error) {
	klog.Infoln("Called EKSDescribeNodegroup")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}
	out, err := eksSvc.DescribeNodegroup(&Input.EKSDescribeNodegroupInput)
	return out, err
}

func EKSListNodegroup(Input util.HCPListNodegroupInput) (*eks.ListNodegroupsOutput, error) {
	klog.Infoln("Called EKSListNodegroup")

	eksSvc := InitializeEKSClient(Input.Region)
	if eksSvc == nil {
		return nil, nil
	}
	out, err := eksSvc.ListNodegroups(&Input.EKSListNodegroupInput)
	return out, err
}

func EKSCreateAddon(addonInput eks.CreateAddonInput) (*eks.CreateAddonOutput, error) {

	klog.Info("Called EKSCreateAddon")
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

func EKSDeleteAddon(addonInput eks.DeleteAddonInput) (*eks.DeleteAddonOutput, error) {

	klog.Info("Called EKSDeleteAddon")
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

func EKSDescribeAddon(addonInput eks.DescribeAddonInput) (*eks.DescribeAddonOutput, error) {

	klog.Info("Called EKSDescribeAddon")
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

func EKSDescribeAddonVersions(addonInput eks.DescribeAddonVersionsInput) (*eks.DescribeAddonVersionsOutput, error) {

	klog.Info("Called EKSDescribeAddonVersions")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)

	newAddonInput := &eks.DescribeAddonVersionsInput{
		AddonName:         addonInput.AddonName,
		KubernetesVersion: addonInput.KubernetesVersion,
		MaxResults:        addonInput.MaxResults,
		NextToken:         addonInput.NextToken,
	}
	klog.Info(&newAddonInput.MaxResults)
	out, err := eksSvc.DescribeAddonVersions(newAddonInput)

	return out, err
}

func EKSListAddon(addonInput eks.ListAddonsInput) (*eks.ListAddonsOutput, error) {

	klog.Info("Called EKSListAddon")
	eksSvc, err := GetEKSClient(addonInput.ClusterName)
	if eksSvc == nil {
		return nil, err
	}

	newAddonInput := &eks.ListAddonsInput{
		ClusterName: addonInput.ClusterName,
		MaxResults:  addonInput.MaxResults,
		NextToken:   addonInput.NextToken,
	}
	out, err := eksSvc.ListAddons(newAddonInput)

	return out, err
}

func EKSUpdateAddon(addonInput eks.UpdateAddonInput) (*eks.UpdateAddonOutput, error) {

	klog.Info("Called EKSUpdateAddon")
	eksSvc, err := GetEKSClient(addonInput.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newAddonInput := &eks.UpdateAddonInput{
		ClusterName:           addonInput.ClusterName,
		AddonName:             addonInput.AddonName,
		AddonVersion:          addonInput.AddonVersion,
		ServiceAccountRoleArn: addonInput.ServiceAccountRoleArn,
		ResolveConflicts:      addonInput.ResolveConflicts,
		ClientRequestToken:    addonInput.ClientRequestToken,
	}
	out, err := eksSvc.UpdateAddon(newAddonInput)

	return out, err
}

func EKSAssociateEncryptionConfig(input eks.AssociateEncryptionConfigInput) (*eks.AssociateEncryptionConfigOutput, error) {

	klog.Info("Called EKSAssociateEncryptionConfig")
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

func EKSAssociateIdentityProviderConfig(input eks.AssociateIdentityProviderConfigInput) (*eks.AssociateIdentityProviderConfigOutput, error) {

	klog.Info("Called EKSAssociateIdentityProviderConfig")
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

func EKSDisassociateIdentityProviderConfig(input eks.DisassociateIdentityProviderConfigInput) (*eks.DisassociateIdentityProviderConfigOutput, error) {

	klog.Info("Called EKSDisassociateIdentityProviderConfig")
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

func EKSDescribeIdentityProviderConfig(input eks.DescribeIdentityProviderConfigInput) (*eks.DescribeIdentityProviderConfigOutput, error) {

	klog.Info("Called EKSDescribeIdentityProviderConfig")
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

func EKSListIdentityProviderConfigs(input eks.ListIdentityProviderConfigsInput) (*eks.ListIdentityProviderConfigsOutput, error) {

	klog.Info("Called EKSListIdentityProviderConfigs")
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

func EKSListTagsForResource(listTagsForResourceInput eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {

	klog.Info("Called EKSListTagsForResource")
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

func EKSTagResource(input eks.TagResourceInput) (*eks.TagResourceOutput, error) {

	klog.Info("Called EKSTagResource")
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

func EKSUntagResource(input eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {

	klog.Info("Called EKSUntagResource")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	eksSvc := eks.New(sess)
	input = eks.UntagResourceInput{
		ResourceArn: input.ResourceArn,
		TagKeys:     input.TagKeys,
	}
	out, err := eksSvc.UntagResource(&input)

	return out, err
}

// update

func EKSListUpdate(listUpdateInput eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {

	klog.Info("Called EKSListUpdate")
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

func EKSDescribeUpdate(describeUpdateInput eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {

	klog.Info("Called EKSDescribeUpdate")
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

func EKSUpdateClusterConfig(input eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error) {

	klog.Info("Called EKSUpdateClusterConfig")
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

func EKSUpdateNodeGroupConfig(input eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {

	klog.Info("Called EKSUpdateNodeGroupConfig")
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
