package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/KETI-Hybrid/hcp-analytic-engine-v1/influx"

	"github.com/influxdata/influxdb/client/v2"
	"k8s.io/klog"
)

type PodMetricList struct {
	Items []PodMetric `json:"podmetrics"`
}
type NodeMetricList struct {
	Items []NodeMetric `json:"nodemetrics"`
}
type PodMetric struct {
	Time      string        `json:"time"`
	Cluster   string        `json:"cluster"`
	Namespace string        `json:"namespace"`
	Node      string        `json:"node"`
	Pod       string        `json:"pod"`
	Cpu       CpuMetric     `json:"cpu"`
	Memory    MemoryMetric  `json:"memory"`
	Fs        FsMetric      `json:"fs"`
	Network   NetworkMetric `json:"network"`
}
type NodeMetric struct {
	Time    string        `json:"time"`
	Cluster string        `json:"cluster"`
	Node    string        `json:"node"`
	Cpu     CpuMetric     `json:"cpu"`
	Memory  MemoryMetric  `json:"memory"`
	Fs      FsMetric      `json:"fs"`
	Network NetworkMetric `json:"network"`
}

type CpuMetric struct {
	CPUUsageNanoCores string `json:"CPUUsageNanoCores"`
}
type MemoryMetric struct {
	MemoryAvailableBytes  string `json:"MemoryAvailableBytes"`
	MemoryUsageBytes      string `json:"MemoryUsageBytes"`
	MemoryWorkingSetBytes string `json:"MemoryWorkingSetBytes"`
}
type FsMetric struct {
	FsAvailableBytes string `json:"FsAvailableBytes"`
	FsCapacityBytes  string `json:"FsCapacityBytes"`
	FsUsedBytes      string `json:"FsUsedBytes"`
}
type NetworkMetric struct {
	NetworkRxBytes string `json:"NetworkRxBytes"`
	NetworkTxBytes string `json:"NetworkTxBytes"`
}

func GetNodeMetric(w http.ResponseWriter, r *http.Request) {

}

func GetResource(podNum int, clusterName string, objectType string) []byte {
	INFLUX_IP := os.Getenv("INFLUX_IP")
	INFLUX_PORT := os.Getenv("INFLUX_PORT")
	INFLUX_USERNAME := os.Getenv("INFLUX_USERNAME")
	INFLUX_PASSWORD := os.Getenv("INFLUX_PASSWORD")

	inf := influx.NewInflux(INFLUX_IP, INFLUX_PORT, INFLUX_USERNAME, INFLUX_PASSWORD)

	if objectType == "pods" {

		results := inf.GetPodData(podNum, clusterName)

		pm := setPodMetric(results)

		bytesJson, _ := json.Marshal(pm)
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, bytesJson, "", "\t")
		if err != nil {
			panic(err.Error())
		}

		return prettyJSON.Bytes()

	} else if objectType == "nodes" {
		results := inf.GetNodeData(clusterName)
		nm := setNodeMetric(results)

		bytesJson, _ := json.Marshal(nm)
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, bytesJson, "", "\t")
		if err != nil {
			panic(err.Error())
		}

		return prettyJSON.Bytes()

	} else {
		klog.V(0).Info("Error : objectType is only pods or nodes")
		return nil
	}

}

func setPodMetric(results []client.Result) *PodMetricList {
	pmList := &PodMetricList{}
	for _, result := range results {
		for _, ser := range result.Series {

			for r, _ := range ser.Values {
				pm := &PodMetric{}
				for c, colName := range ser.Columns {

					value := ser.Values[r][c]
					if colName == "time" {
						pm.Time = value.(string)
					} else if colName == "CPUUsageNanoCores" {
						pm.Cpu.CPUUsageNanoCores = value.(string)
					} else if colName == "FsAvailableBytes" {
						pm.Fs.FsAvailableBytes = value.(string)
					} else if colName == "FsCapacityBytes" {
						pm.Fs.FsCapacityBytes = value.(string)
					} else if colName == "FsUsedBytes" {
						pm.Fs.FsUsedBytes = value.(string)
					} else if colName == "MemoryAvailableBytes" {
						pm.Memory.MemoryAvailableBytes = value.(string)
					} else if colName == "MemoryUsageBytes" {
						pm.Memory.MemoryUsageBytes = value.(string)
					} else if colName == "MemoryWorkingSetBytes" {
						pm.Memory.MemoryWorkingSetBytes = value.(string)
					} else if colName == "NetworkRxBytes" {
						pm.Network.NetworkRxBytes = value.(string)
					} else if colName == "NetworkTxBytes" {
						pm.Network.NetworkTxBytes = value.(string)
					} else if colName == "cluster" {
						pm.Cluster = value.(string)
					} else if colName == "namespace" {
						pm.Namespace = value.(string)
					} else if colName == "node" {
						pm.Node = value.(string)
					} else if colName == "pod" {
						pm.Pod = value.(string)
					}

				}
				pmList.Items = append(pmList.Items, *pm)
			}

		}
	}

	return pmList

}
func setNodeMetric(results []client.Result) *NodeMetricList {
	nmList := &NodeMetricList{}
	for _, result := range results {
		for _, ser := range result.Series {
			for r, _ := range ser.Values {
				nm := &NodeMetric{}
				for c, colName := range ser.Columns {
					value := ser.Values[r][c]
					if colName == "time" {
						nm.Time = value.(string)
					} else if colName == "CPUUsageNanoCores" {
						nm.Cpu.CPUUsageNanoCores = value.(string)
					} else if colName == "FsAvailableBytes" {
						nm.Fs.FsAvailableBytes = value.(string)
					} else if colName == "FsCapacityBytes" {
						nm.Fs.FsCapacityBytes = value.(string)
					} else if colName == "FsUsedBytes" {
						nm.Fs.FsUsedBytes = value.(string)
					} else if colName == "MemoryAvailableBytes" {
						nm.Memory.MemoryAvailableBytes = value.(string)
					} else if colName == "MemoryUsageBytes" {
						nm.Memory.MemoryUsageBytes = value.(string)
					} else if colName == "MemoryWorkingSetBytes" {
						nm.Memory.MemoryWorkingSetBytes = value.(string)
					} else if colName == "NetworkRxBytes" {
						nm.Network.NetworkRxBytes = value.(string)
					} else if colName == "NetworkTxBytes" {
						nm.Network.NetworkTxBytes = value.(string)
					} else if colName == "cluster" {
						nm.Cluster = value.(string)
					} else if colName == "node" {
						nm.Node = value.(string)
					}

				}
				nmList.Items = append(nmList.Items, *nm)
			}

		}

	}

	return nmList

}
