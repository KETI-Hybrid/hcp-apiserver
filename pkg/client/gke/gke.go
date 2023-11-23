package gke

import (
	"context"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/option"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

type GKEClient struct {
	Token           string
	ProjectID       string
	ClusterName     string
	ContanerService *container.Service
}

func NewGKEClient(k8sclient *kubernetes.Clientset) *GKEClient {
	client := &GKEClient{}
	config, err := k8sclient.CoreV1().ConfigMaps("public-auth").Get(context.Background(), "gke-auth", metav1.GetOptions{})
	if err != nil {
		klog.Errorln(err.Error())
	}
	
	client.ProjectID = config.Data["project-id"]
	client.ClusterName = config.Data["cluster-name"]
	client.Token = config.Data["key.json"]
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	client.ContanerService, err = container.NewService(ctx, option.WithHTTPClient(c))
	if err != nil {
		log.Fatal(err)
	}
	
	return client
}
