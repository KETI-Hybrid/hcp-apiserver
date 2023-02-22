package gke

import (
	"context"
	"io"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

type GKEClient struct {
	Client      *http.Client
	Request     *http.Request
	Token       string
	Zone        string
	ProjectID   string
	ClusterName string
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
	client.Zone = config.Data["zone"]
	client.Client = &http.Client{}
	return client
}

func (gclient *GKEClient) SetRequestWithBody(url string, body io.Reader) {
	var err error
	gclient.Request, err = http.NewRequest("GET", url, body)
	if err != nil {
		klog.Errorln(err)
	}

	// HTTP 요청에 필요한 인증 정보를 설정합니다.
	gclient.Request.Header.Set("Authorization", "Bearer "+gclient.Token)
}

func (gclient *GKEClient) SetRequestWithoutBody(url string) {
	var err error
	gclient.Request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		klog.Errorln(err)
	}

	// HTTP 요청에 필요한 인증 정보를 설정합니다.
	gclient.Request.Header.Set("Authorization", "Bearer "+gclient.Token)
}

func (gclient *GKEClient) Do() *http.Response {
	resp, err := gclient.Client.Do(gclient.Request)
	if err != nil {
		klog.Errorln(err)
	}
	defer resp.Body.Close()
	return resp
}
