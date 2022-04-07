package gke

import (
	"Hybrid_Cloud/hcp-apiserver/pkg/util"
	hybridctlutil "Hybrid_Cloud/hybridctl/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	container "cloud.google.com/go/container/apiv1"
	"google.golang.org/api/option"
	"google.golang.org/api/sourcerepo/v1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

// gke container images

func ImagesAddTag(w http.ResponseWriter, req *http.Request) {
	var i hybridctlutil.GKEImages
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
	var i hybridctlutil.GKEImages
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
	var i hybridctlutil.GKEImages
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
	var i hybridctlutil.GKEImages
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
	var i *hybridctlutil.GKEImages
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
	var i *hybridctlutil.GKEImages
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
	var op *hybridctlutil.GKEOperations
	SetGKERequest(r, &op)
	cmd := exec.Command("gcloud", "container", "operations", "wait", op.OPERATION_ID)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

// gcloud auth configure-docker
func ConfigureDocker(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("gcloud", "auth", "configure-docker")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func AuthList(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("gcloud", "auth", "list")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func AuthRevoke(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("gcloud", "auth", "revoke")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var a *hybridctlutil.GKEAuth
	util.Parser(r, a)
	w.Header().Set("Content-Type", "application/json")
	var data []byte
	var err error
	var str string = "You are already authenticated with 'hybridcloudplatform@keti-container.iam.gserviceaccount.com'.\nDo you wish to proceed and overwrite existing credentials?\n\nDo you want to continue (Y/n)?"
	if a.CRED_FILE == "" {
		str := "ERROR: Input path to the external account configuration file "
		data, err = json.Marshal(str)
	} else {
		cmd := exec.Command("gcloud", "auth", "login", "--cred-file", a.CRED_FILE)
		data, err = util.GetOutputReplaceStr(cmd, str, "")
	}

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
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

func UpdateProjectConfig(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	sourcerepoService, err := sourcerepo.NewService(ctx)
	if err != nil {
		fmt.Println(err)
	}
	projectsService := sourcerepo.NewProjectsService(sourcerepoService)
	var req *sourcerepo.UpdateProjectConfigRequest
	SetGKERequest(r, &req)

	call := projectsService.UpdateConfig("", req)
	resp, err := call.Do()

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

func GetProjectConfig(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	sourcerepoService, err := sourcerepo.NewService(ctx)
	if err != nil {
		fmt.Println(err)
	}
	projectsService := sourcerepo.NewProjectsService(sourcerepoService)
	var req *sourcerepo.UpdateProjectConfigRequest
	SetGKERequest(r, &req)

	call := projectsService.GetConfig("keti-container")
	resp, err := call.Do()

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

func ConfigSet(w http.ResponseWriter, r *http.Request) {
	var s *hybridctlutil.GKESetProperty
	util.Parser(r, s)
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
