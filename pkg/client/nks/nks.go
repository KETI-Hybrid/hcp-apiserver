package nks

import (
	"context"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/vnks"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

type NKSClient struct {
	Client    *vnks.APIClient
	AccessKey string
	SecretKey string
}

func NewNKSClient(k8sclient *kubernetes.Clientset) *NKSClient {
	client := &NKSClient{}
	config, err := k8sclient.CoreV1().ConfigMaps("public-auth").Get(context.Background(), "nks-auth", metav1.GetOptions{})
	if err != nil {
		klog.Errorln(err.Error())
	}
	apiKeys := ncloud.Keys()
	apiKeys.AccessKey = config.Data["accessKey"]
	client.AccessKey = config.Data["accessKey"]
	apiKeys.SecretKey = config.Data["secretKey"]
	client.SecretKey = config.Data["secretKey"]
	client.Client = vnks.NewAPIClient(vnks.NewConfiguration(config.Data["region"], apiKeys))

	return client
}
