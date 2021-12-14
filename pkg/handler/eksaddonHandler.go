package handler

import (
	clusterRegister "Hybrid_Cluster/pkg/client/clusterregister/v1alpha1/clientset/versioned/typed/clusterregister/v1alpha1"
	"context"

	// /root/Go/src/Hybrid_Cluster/pkg/client/clientset/v1alpha1
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/eks"

	cobrautil "Hybrid_Cluster/hybridctl/util"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

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
