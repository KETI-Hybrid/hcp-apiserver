package main

/*
import (
	"hcp-apiserver/pkg/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AKS_Cluster_API struct {
	Value []struct {
		ID       string `json:"id"`
		Location string `json:"location"`
		Name     string `json:"name"`
		Tags     struct {
			Environment string `json:"Environment"`
		} `json:"tags"`
		Type       string `json:"type"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
			PowerState        struct {
				Code string `json:"code"`
			} `json:"powerState"`
			KubernetesVersion string `json:"kubernetesVersion"`
			DNSPrefix         string `json:"dnsPrefix"`
			Fqdn              string `json:"fqdn"`
			AzurePortalFQDN   string `json:"azurePortalFQDN"`
			AgentPoolProfiles []struct {
				Name              string `json:"name"`
				Count             int    `json:"count"`
				VMSize            string `json:"vmSize"`
				OsDiskSizeGB      int    `json:"osDiskSizeGB"`
				OsDiskType        string `json:"osDiskType"`
				KubeletDiskType   string `json:"kubeletDiskType"`
				MaxPods           int    `json:"maxPods"`
				Type              string `json:"type"`
				EnableAutoScaling bool   `json:"enableAutoScaling"`
				ProvisioningState string `json:"provisioningState"`
				PowerState        struct {
					Code string `json:"code"`
				} `json:"powerState"`
				OrchestratorVersion    string `json:"orchestratorVersion"`
				EnableNodePublicIP     bool   `json:"enableNodePublicIP"`
				Mode                   string `json:"mode"`
				EnableEncryptionAtHost bool   `json:"enableEncryptionAtHost"`
				OsType                 string `json:"osType"`
				OsSKU                  string `json:"osSKU"`
				NodeImageVersion       string `json:"nodeImageVersion"`
				EnableFIPS             bool   `json:"enableFIPS"`
			} `json:"agentPoolProfiles"`
			LinuxProfile struct {
				AdminUsername string `json:"adminUsername"`
				SSH           struct {
					PublicKeys []struct {
						KeyData string `json:"keyData"`
					} `json:"publicKeys"`
				} `json:"ssh"`
			} `json:"linuxProfile"`
			ServicePrincipalProfile struct {
				ClientID string `json:"clientId"`
			} `json:"servicePrincipalProfile"`
			AddonProfiles struct {
				Omsagent struct {
					Enabled bool `json:"enabled"`
					Config  struct {
						LogAnalyticsWorkspaceResourceID string `json:"logAnalyticsWorkspaceResourceID"`
					} `json:"config"`
				} `json:"omsagent"`
			} `json:"addonProfiles"`
			NodeResourceGroup string `json:"nodeResourceGroup"`
			EnableRBAC        bool   `json:"enableRBAC"`
			NetworkProfile    struct {
				NetworkPlugin       string `json:"networkPlugin"`
				LoadBalancerSku     string `json:"loadBalancerSku"`
				LoadBalancerProfile struct {
					ManagedOutboundIPs struct {
						Count int `json:"count"`
					} `json:"managedOutboundIPs"`
					EffectiveOutboundIPs []struct {
						ID string `json:"id"`
					} `json:"effectiveOutboundIPs"`
				} `json:"loadBalancerProfile"`
				PodCidr          string `json:"podCidr"`
				ServiceCidr      string `json:"serviceCidr"`
				DNSServiceIP     string `json:"dnsServiceIP"`
				DockerBridgeCidr string `json:"dockerBridgeCidr"`
				OutboundType     string `json:"outboundType"`
			} `json:"networkProfile"`
			MaxAgentPools          int `json:"maxAgentPools"`
			APIServerAccessProfile struct {
				EnablePrivateCluster bool `json:"enablePrivateCluster"`
			} `json:"apiServerAccessProfile"`
		} `json:"properties"`
		Sku struct {
			Name string `json:"name"`
			Tier string `json:"tier"`
		} `json:"sku"`
	} `json:"value"`
}

func main() {
	url := "https://management.azure.com/subscriptions/{subscriptionsId}/providers/Microsoft.ContainerService/managedClusters?api-version=2021-05-01"
	resp, _ := util.AuthorizationAndHTTP("GET", url, nil)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	// string to struck 과정 필요
	fmt.Println("str: ", str)
	data := AKS_Cluster_API{} //json to struck
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}
	fmt.Println("--- Show aks Cluster list ---")
	// 출력
	// fmt.Println(data.Value)
	for i := 0; i < len(data.Value); i++ {
		fmt.Println("[", i+1, "]", data.Value[i].Name)
	}
}
*/
