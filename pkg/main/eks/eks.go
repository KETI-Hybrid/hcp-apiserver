package eks

import (
	"Hybrid_Cluster/hcp-apiserver/pkg/handler"
	"Hybrid_Cluster/hcp-apiserver/pkg/util"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/service/eks"
)

func CreateAddon(w http.ResponseWriter, req *http.Request) {

	var createAddonInput eks.CreateAddonInput

	util.Parser(w, req, &createAddonInput)
	out, err := handler.CreateAddon(createAddonInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func DeleteAddon(w http.ResponseWriter, req *http.Request) {

	var deleteAddonInput eks.DeleteAddonInput

	util.Parser(w, req, &deleteAddonInput)
	out, err := handler.DeleteAddon(deleteAddonInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func DescribeAddon(w http.ResponseWriter, req *http.Request) {

	var describeAddonInput eks.DescribeAddonInput

	util.Parser(w, req, &describeAddonInput)
	out, err := handler.DescribeAddon(describeAddonInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func DescribeAddonVersions(w http.ResponseWriter, req *http.Request) {

	var describeAddonVersionsInput eks.DescribeAddonVersionsInput

	util.Parser(w, req, &describeAddonVersionsInput)
	out, err := handler.DescribeAddonVersions(describeAddonVersionsInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))

}

func ListAddon(w http.ResponseWriter, req *http.Request) {

	var listAddonInput eks.ListAddonsInput

	util.Parser(w, req, &listAddonInput)
	out, err := handler.ListAddon(listAddonInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))

}

func UpdateAddon(w http.ResponseWriter, req *http.Request) {

	var updateAddonInput eks.UpdateAddonInput

	util.Parser(w, req, &updateAddonInput)
	out, err := handler.UpdateAddon(updateAddonInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func ListUpdate(w http.ResponseWriter, req *http.Request) {
	var listUpdateInput eks.ListUpdatesInput

	util.Parser(w, req, &listUpdateInput)
	out, err := handler.ListUpdate(listUpdateInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func DescribeUpdate(w http.ResponseWriter, req *http.Request) {
	var describeUpdateInput eks.DescribeUpdateInput

	util.Parser(w, req, &describeUpdateInput)
	out, err := handler.DescribeUpdate(describeUpdateInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func ListTagsForResource(w http.ResponseWriter, req *http.Request) {
	var listTagsForResourceInput eks.ListTagsForResourceInput

	util.Parser(w, req, &listTagsForResourceInput)
	out, err := handler.ListTagsForResource(listTagsForResourceInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func AssociateIdentityProviderConfig(w http.ResponseWriter, req *http.Request) {
	var associateIdentityProviderConfigInput eks.AssociateIdentityProviderConfigInput

	util.Parser(w, req, &associateIdentityProviderConfigInput)
	out, err := handler.AssociateIdentityProviderConfig(associateIdentityProviderConfigInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func AssociateEncryptionConfig(w http.ResponseWriter, req *http.Request) {
	var associateEncryptionConfigInput eks.AssociateEncryptionConfigInput

	util.Parser(w, req, &associateEncryptionConfigInput)
	out, err := handler.AssociateEncryptionConfig(associateEncryptionConfigInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func DisassociateIdentityProviderConfig(w http.ResponseWriter, req *http.Request) {
	var disassociateIdentityProviderConfigInput eks.DisassociateIdentityProviderConfigInput

	util.Parser(w, req, &disassociateIdentityProviderConfigInput)
	out, err := handler.DisassociateIdentityProviderConfig(disassociateIdentityProviderConfigInput)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func DescribeIdentityProviderConfig(w http.ResponseWriter, req *http.Request) {
	var input eks.DescribeIdentityProviderConfigInput

	util.Parser(w, req, &input)
	out, err := handler.DescribeIdentityProviderConfig(input)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func ListIdentityProviderConfigs(w http.ResponseWriter, req *http.Request) {
	var input eks.ListIdentityProviderConfigsInput

	util.Parser(w, req, &input)
	out, err := handler.ListIdentityProviderConfigs(input)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func TagResource(w http.ResponseWriter, req *http.Request) {
	var input eks.TagResourceInput

	util.Parser(w, req, &input)
	out, err := handler.TagResource(input)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func UntagResource(w http.ResponseWriter, req *http.Request) {
	var input eks.UntagResourceInput

	util.Parser(w, req, &input)
	out, err := handler.UntagResource(input)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func UpdateClusterConfig(w http.ResponseWriter, req *http.Request) {
	var input eks.UpdateClusterConfigInput

	util.Parser(w, req, &input)
	out, err := handler.UpdateClusterConfig(input)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}

func UpdateNodegroupConfig(w http.ResponseWriter, req *http.Request) {
	var input eks.UpdateNodegroupConfigInput

	util.Parser(w, req, &input)
	out, err := handler.UpdateNodeGroupConfig(input)
	var jsonData []byte
	if err != nil {
		jsonData, _ = json.Marshal(&err)
	} else {
		jsonData, _ = json.Marshal(&out)
	}
	w.Write([]byte(jsonData))
}
