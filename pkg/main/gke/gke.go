package gke

import (
	"Hybrid_Cloud/hcp-apiserver/pkg/util"
	cobrautil "Hybrid_Cloud/hybridctl/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	container "cloud.google.com/go/container/apiv1"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

// gke container images

func ImagesAddTag(w http.ResponseWriter, req *http.Request) {
	var i cobrautil.GKEImages
	util.Parser(req, &i)

	args := []string{"container", "images", "add-tag", i.SRC_IMAGE}
	args = append(args, i.DEST_IMAGE...)

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")

	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func ImagesDelete(w http.ResponseWriter, req *http.Request) {
	var i cobrautil.GKEImages
	util.Parser(req, &i)

	args := []string{"container", "images", "delete"}
	args = append(args, i.IMAGE_NAME...)

	if i.FORCE_DELETE_TAGS {
		args = append(args, "--force-delete-tags")
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")

	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func ImagesDescribe(w http.ResponseWriter, req *http.Request) {
	var i cobrautil.GKEImages
	util.Parser(req, &i)

	cmd := exec.Command("gcloud", "container", "images", "describe", i.IMAGE_NAME[0])
	data, err := util.GetOutput(cmd)

	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func ImagesList(w http.ResponseWriter, req *http.Request) {
	var i cobrautil.GKEImages
	util.Parser(req, &i)

	args := []string{"container", "images", "list"}
	if i.REPOSITORY != "" {
		args = append(args, "--repository", i.REPOSITORY)
	}

	if i.FILTER != "" {
		args = append(args, "--filter", i.FILTER)
	}

	if i.LIMIT != "" {
		args = append(args, "--limit", i.LIMIT)
	}

	if i.PAGE_SIZE != "" {
		args = append(args, "--page-size", i.PAGE_SIZE)
	}

	if i.SORT_BY != "" {
		args = append(args, "--sort-by", i.SORT_BY)
	}

	if i.URI {
		args = append(args, "--uri")
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)

	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func ImagesListTags(w http.ResponseWriter, req *http.Request) {
	var i *cobrautil.GKEImages
	util.Parser(req, &i)

	args := []string{"container", "images", "list-tags", i.IMAGE_NAME[0]}
	if i.FILTER != "" {
		args = append(args, "--filter", i.FILTER)
	}

	if i.LIMIT != "" {
		args = append(args, "--limit", i.LIMIT)
	}

	if i.PAGE_SIZE != "" {
		args = append(args, "--page-size", i.PAGE_SIZE)
	}

	if i.SORT_BY != "" {
		args = append(args, "--sort-by", i.SORT_BY)
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)

	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func ImagesUnTags(w http.ResponseWriter, req *http.Request) {
	var i *cobrautil.GKEImages
	util.Parser(req, &i)

	args := []string{"container", "images", "untag"}
	args = append(args, i.IMAGE_NAME...)
	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")

	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// create NewClusterManagerClient
func NewClusterManagerClient() (*container.ClusterManagerClient, error) {
	ctx := context.TODO()
	c, err := container.NewClusterManagerClient(ctx, option.WithCredentialsFile("/root/hcp-key.json"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//defer c.Close()
	return c, err
}

// unmarshalling request to Operation struct
func SetGKERequest(r *http.Request, input interface{}) {
	jsonDataFromHttp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err = json.Unmarshal(jsonDataFromHttp, &input); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(input)
	}
}

func GetServerConfig(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	var req containerpb.GetServerConfigRequest
	SetGKERequest(r, &req)
	fmt.Println(req)

	resp, err := c.GetServerConfig(context.TODO(), &req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

// https://pkg.go.dev/cloud.google.com/go/container/apiv1
func GetOperation(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()

	if err != nil {
		fmt.Println(err)
	}

	var req *containerpb.GetOperationRequest
	SetGKERequest(r, &req)

	resp, err := c.GetOperation(context.TODO(), req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

func ListOperations(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	var req *containerpb.ListOperationsRequest
	SetGKERequest(r, &req)

	resp, err := c.ListOperations(context.TODO(), req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

func RollbackNodePoolUpgrade(w http.ResponseWriter, r *http.Request) {
	c, err := NewClusterManagerClient()
	if err != nil {
		fmt.Println(err)
	}

	var req *containerpb.RollbackNodePoolUpgradeRequest
	SetGKERequest(r, &req)

	resp, err := c.RollbackNodePoolUpgrade(context.TODO(), req)
	defer c.Close()

	var output util.Output
	if err != nil {
		bytes, _ := json.Marshal(err.Error())
		output.Stderr = bytes
	} else {
		bytes, _ := json.Marshal(&resp)
		output.Stdout = bytes
	}

	bytes, _ := json.Marshal(output)
	w.Write(bytes)
}

func WaitOperations(w http.ResponseWriter, r *http.Request) {
	var input cobrautil.GKEOperations
	SetGKERequest(r, &input)

	args := []string{"container", "operations", "wait", input.OPERATION_ID}
	if input.ZONE != "" {
		args = append(args, "-z", input.ZONE)
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

// gcloud auth configure-docker
func AuthConfigureDocker(w http.ResponseWriter, r *http.Request) {
	var input cobrautil.GKEAuth
	SetGKERequest(r, &input)

	fmt.Println(input.REGISTRIES)
	args := []string{"auth", "configure-docker", input.REGISTRIES}

	/*
		args = append(args, input.REGISTRIES...)
	*/

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func AuthList(w http.ResponseWriter, r *http.Request) {
	var input cobrautil.GKEAuth
	SetGKERequest(r, &input)

	args := []string{"auth", "list"}
	if input.FILTER_ACCOUNT != "" {
		args = append(args, "--filter-account", input.FILTER_ACCOUNT)
	}

	if input.FILTER != "" {
		args = append(args, "--filter", input.FILTER)
	}

	if input.LIMIT != "" {
		args = append(args, "--limit", input.LIMIT)
	}

	if input.PAGE_SIZE != "" {
		args = append(args, "--page-size", input.PAGE_SIZE)
	}

	if input.SORT_BY != "" {
		args = append(args, "--sort-by", input.SORT_BY)
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func AuthRevoke(w http.ResponseWriter, r *http.Request) {
	var input cobrautil.GKEAuth
	SetGKERequest(r, &input)

	args := []string{"auth", "revoke"}
	if input.ACCOUNTS != "" {
		args = append(args, input.ACCOUNTS)
	}

	if input.ALL {
		args = append(args, "--all")
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var input cobrautil.GKEAuth
	util.Parser(r, &input)

	args := []string{"auth", "login"}
	if input.ACCOUNTS != "" {
		args = append(args, input.ACCOUNTS)
	}

	args = append(args, "--cred-file", input.CRED_FILE)
	var str string = "Do you wish to proceed and overwrite existing credentials?\n\nDo you want to continue (Y/n)?"
	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutputReplaceStr(cmd, str, "")

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// gcloud docker
type Docker struct {
	AUTHORIZE_ONLY bool
	DOCKER_HOST    string
	SERVER         string
}

func GDocker(w http.ResponseWriter, req *http.Request) {
	var d *Docker
	util.Parser(req, d)
	w.Header().Set("Content-Type", "application/json")
	args := []string{"docker"}
	if d.AUTHORIZE_ONLY {
		args = append(args, "-a")
	}

	if d.DOCKER_HOST != "" {
		args = append(args, "--docker-host", d.DOCKER_HOST)
	}

	if d.SERVER != "" {
		args = append(args, "-s", d.SERVER)
	}

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func UpdateProjectConfigs(w http.ResponseWriter, r *http.Request) {
	var input cobrautil.GKESource
	util.Parser(r, &input)

	args := []string{"source", "project-configs", "update"}
	if input.PUSHBLOCK == 0 {
		args = append(args, "--disable-pushblock")
	} else if input.PUSHBLOCK == 1 {
		args = append(args, "--enable-pushblock")
	}

	if input.MESSAGE_FORMAT != "" {
		args = append(args, "--message-format", input.MESSAGE_FORMAT)
	}

	if input.SERVICE_ACCOUNT != "" {
		args = append(args, "--service-account", input.SERVICE_ACCOUNT)
	}

	if input.TOPIC_PROJECT != "" {
		args = append(args, "--topic-project", input.TOPIC_PROJECT)
	}

	if input.ADD_TOPIC != "" {
		args = append(args, "--add-topic", input.ADD_TOPIC)
	}

	if input.REMOVE_TOPIC != "" {
		args = append(args, "--remove-topic", input.REMOVE_TOPIC)
	}

	if input.UPDATE_TOPIC != "" {
		args = append(args, "--update-topic", input.UPDATE_TOPIC)
	}

	cmd := exec.Command("gcloud", args...)
	fmt.Println(cmd.Args)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

}

func DescribeProjectConfigs(w http.ResponseWriter, r *http.Request) {
	var input cobrautil.GKESource
	util.Parser(r, &input)

	args := []string{"source", "project-configs", "describe"}

	cmd := exec.Command("gcloud", args...)
	fmt.Println(cmd.Args)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

}

func ConfigSet(w http.ResponseWriter, r *http.Request) {
	var s cobrautil.GKESetProperty
	util.Parser(r, &s)
	w.Header().Set("Content-Type", "application/json")
	args := []string{"config", "set"}

	// SECTION/ is optional while referring to properties in the core section
	if s.SECTION != "" {
		if s.PROPERTY != "" {
			// gcloud config set SECTION/PROPERTY VALUE
			str := s.SECTION + "/" + s.PROPERTY
			args = append(args, str)
			if s.VALUE != "" {
				args = append(args, s.VALUE)
			}
		}
	} else {
		// gcloud config set SECTION/PROPERTY VALUE
		if s.PROPERTY != "" {
			args = append(args, s.PROPERTY)
			if s.VALUE != "" {
				args = append(args, s.VALUE)
			}
		}
	}

	if s.INSTALLATION {
		args = append(args, "--installation")
	}
	fmt.Println(args)

	cmd := exec.Command("gcloud", args...)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}
