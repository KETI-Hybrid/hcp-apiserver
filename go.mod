module hcp-apiserver

go 1.18

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.8
	k8s.io/client-go => k8s.io/client-go v0.22.8
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
)

require github.com/julienschmidt/httprouter v1.3.0
