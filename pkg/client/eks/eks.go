package eks

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

func NewEKSClient(k8sclient *kubernetes.Clientset) *session.Session {
	config, err := k8sclient.CoreV1().ConfigMaps("public-auth").Get(context.Background(), "aws-auth", metav1.GetOptions{})
	if err != nil {
		klog.Errorln(err.Error())
	}

	region := config.Data["region"]
	id := config.Data["aws_access_key_id"]
	secret := config.Data["aws_secret_access_key"]
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			id,
			secret,
			""),
	})

	if err != nil {
		klog.Errorln(err)
	}
	return sess
}
