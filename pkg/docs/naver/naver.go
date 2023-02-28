package naver

import (
	"hcp-apiserver/pkg/apis/naver/cluster"
	"hcp-apiserver/pkg/apis/naver/nodepools"
	"hcp-apiserver/pkg/apis/naver/workernode"

	"github.com/julienschmidt/httprouter"
)

func InitNKSEndPoint(router *httprouter.Router) {
	cluster.ClusterResourceAttach(router)
	workernode.WorkerNodeResourceAttach(router)
	nodepools.NodePoolResourceAttach(router)
}
