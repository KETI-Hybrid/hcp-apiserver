package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// type AzureAuth struct {
// 	clientId       string
// 	clientSecret   string
// 	subscriptionId string
// 	tenantId       string
// }

// func GetAzureAuth() AzureAuth {
// 	auth := AzureAuth{}
// 	data, err := ioutil.ReadFile("~/.azure/test.auth")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	json.Unmarshal(data, &auth)

// 	return auth
// }

type bearerToken struct {
	Token_type     string `json:"token_type" protobuf:"bytes,1,opt,name=token_type"`
	Expires_in     string `json:"expires_in" protobuf:"bytes,2,opt,name=expires_in"`
	Ext_expires_in string `json:"ext_expires_in" protobuf:"bytes,3,opt,name=ext_expires_in"`
	Expires_on     string `json:"expires_on" protobuf:"bytes,4,opt,name=expires_on"`
	Not_before     string `json:"not_before" protobuf:"bytes,5,opt,name=not_before"`
	Resource       string `json:"resource" protobuf:"bytes,6,opt,name=resource"`
	Access_token   string `json:"access_token" protobuf:"bytes,7,opt,name=access_token"`
}

func GetBearer() bearerToken {

	// azureAuth := GetAzureAuth()

	params := url.Values{}
	params.Add("client_id", os.Getenv("ClientId"))
	params.Add("grant_type", `client_credentials`)
	params.Add("resource", `https://management.azure.com/`)
	params.Add("client_secret", os.Getenv("ClientSecret"))
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://login.microsoftonline.com/"+os.Getenv("TenantId")+"/oauth2/token", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	token := bearerToken{}
	json.Unmarshal(bytes, &token)

	return token
}

func AuthorizationAndHTTP(method string, hosturl string) (*http.Response, error) {
	params := url.Values{}
	params.Add("resource", `https://management.azure.com/`)
	var request *http.Request
	switch method {
	case "POST":
		body := strings.NewReader(params.Encode())
		request, _ = http.NewRequest(method, hosturl, body)
	case "GET":
		request, _ = http.NewRequest(method, hosturl, nil)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Authorization", "Bearer "+GetBearer().Access_token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err2 : ", err)
	} else {
		fmt.Println(response.Status)
	}
	return response, err
}
