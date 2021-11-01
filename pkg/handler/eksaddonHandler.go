package handler

import (
	clusterRegister "Hybrid_Cluster/clientset/v1alpha1"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/eks"

	cobrautil "Hybrid_Cluster/hybridctl/util"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetEKSClient(clusterName *string) *eks.EKS {
	master_config, _ := cobrautil.BuildConfigFromFlags("kube-master", "/root/.kube/config")
	clusterRegisterClientSet, err := clusterRegister.NewForConfig(master_config)
	checkErr(err)
	clusterRegisters, err := clusterRegisterClientSet.ClusterRegister("eks").Get(*clusterName, metav1.GetOptions{})
	checkErr(err)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(clusterRegisters.Spec.Region),
	}))
	eksSvc := eks.New(sess)
	return eksSvc
}

func CreateAddon(addonInput eks.CreateAddonInput) (*eks.CreateAddonOutput, error) {

	// println(*addonInput.ClusterName)
	eksSvc := GetEKSClient(addonInput.ClusterName)
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

	eksSvc := GetEKSClient(addonInput.ClusterName)

	newAddonInput := &eks.DeleteAddonInput{
		AddonName:   addonInput.AddonName,
		ClusterName: addonInput.ClusterName,
	}
	out, err := eksSvc.DeleteAddon(newAddonInput)

	return out, err
}

func DescribeAddon(addonInput eks.DescribeAddonInput) (*eks.DescribeAddonOutput, error) {

	eksSvc := GetEKSClient(addonInput.ClusterName)

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

	eksSvc := GetEKSClient(addonInput.ClusterName)

	newAddonInput := &eks.ListAddonsInput{
		ClusterName: addonInput.ClusterName,
	}
	out, err := eksSvc.ListAddons(newAddonInput)

	return out, err
}

func UpdateAddon(addonInput eks.UpdateAddonInput) (*eks.UpdateAddonOutput, error) {

	eksSvc := GetEKSClient(addonInput.ClusterName)

	newAddonInput := &eks.UpdateAddonInput{
		ClusterName: addonInput.ClusterName,
		AddonName:   addonInput.AddonName,
	}
	out, err := eksSvc.UpdateAddon(newAddonInput)

	return out, err
}
