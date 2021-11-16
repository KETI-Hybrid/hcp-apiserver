package handler

import (
	"github.com/aws/aws-sdk-go/service/eks"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

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

func AssociateEncryptionConfig(input eks.AssociateEncryptionConfigInput) (*eks.AssociateEncryptionConfigOutput, error) {
	eksSvc, err := GetEKSClient(input.ClusterName)
	if eksSvc == nil {
		return nil, err
	}
	newInput := &eks.AssociateEncryptionConfigInput{
		ClientRequestToken: input.ClientRequestToken,
		ClusterName:        input.ClusterName,
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
