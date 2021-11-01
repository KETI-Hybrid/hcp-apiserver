package converter

// func Converter(info handler.Cli) string {
// 	fmt.Println("---converter start---")

// 	// 명령어 구분
// 	switch info.Cmd {
// 	case "join":
// 		fmt.Printf("---converter end---\n")
// 		return joinConverter(info)
// 	}

// 	fmt.Printf("---converter end---\n")
// 	return ""
// }

// kubernetes platform 구분
// func JoinConverter(info mappingTable.ClusterInfo) string {
// 	fmt.Println("---converter start---")

// 	switch info.PlatformName {
// 	case "gke":
// 		fmt.Println("--> mapping GKE API Arguments & send API to Handler")
// 		api := mappingTable.GetInfo(info)
// 		return api
// 	case "aks":
// 		fmt.Println("--> mapping AKS API Arguments & send API to Handler")
// 		api := mappingTable.AksGetCredential(info)
// 		fmt.Println("---return API---")
// 		fmt.Printf("---converter end---\n")
// 		return api
// 	case "eks":
// 		fmt.Println("--> mapping EKS API Arguments & send API to Handler")
// 		api := mappingTable.GetInfo(info)
// 		return api
// 	default:
// 		return ""
// 	}
// }
