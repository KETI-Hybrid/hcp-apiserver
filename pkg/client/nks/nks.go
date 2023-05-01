package nks

import (
	"context"
	"io"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

type NKSClient struct {
	Client    *http.Client
	Request   *http.Request
	APIURL    string
	ClusterID string
	AccessKey string
	SecretKey string
}

func NewNKSClient(k8sclient *kubernetes.Clientset) *NKSClient {
	client := &NKSClient{}
	config, err := k8sclient.CoreV1().ConfigMaps("public-auth").Get(context.Background(), "nks-auth", metav1.GetOptions{})
	if err != nil {
		klog.Errorln(err.Error())
	}
	client.ClusterID = config.Data["cluster-id"]
	client.AccessKey = config.Data["accessKey"]
	client.SecretKey = config.Data["secretKey"]
	client.APIURL = "https://nks.apigw.fin-ntruss.com/nks/"
	client.Client = &http.Client{}
	return client
}

func (gclient *NKSClient) SetRequestWithBody(url string, body io.Reader) {
	var err error

	gclient.Request, err = http.NewRequest("GET", gclient.APIURL+url, body)
	if err != nil {
		klog.Errorln(err)
	}

	// HTTP 요청에 필요한 인증 정보를 설정합니다.
	gclient.Request.Header.Set("Content-Type", "application/json")
	gclient.Request.Header.Set("X-NCP-APIGW-API-KEY-ID", gclient.AccessKey)
	gclient.Request.Header.Set("X-NCP-APIGW-API-KEY", gclient.SecretKey)
}

func (gclient *NKSClient) SetRequestWithoutBody(url string) {
	var err error
	gclient.Request, err = http.NewRequest("GET", gclient.APIURL+url, nil)
	if err != nil {
		klog.Errorln(err)
	}

	// HTTP 요청에 필요한 인증 정보를 설정합니다.
	gclient.Request.Header.Set("Content-Type", "application/json")
	gclient.Request.Header.Set("X-NCP-APIGW-API-KEY-ID", gclient.AccessKey)
	gclient.Request.Header.Set("X-NCP-APIGW-API-KEY", gclient.SecretKey)
}

func (gclient *NKSClient) Do() *http.Response {
	resp, err := gclient.Client.Do(gclient.Request)
	if err != nil {
		klog.Errorln(err)
	}
	defer resp.Body.Close()
	return resp
}
