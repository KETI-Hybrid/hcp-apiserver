package main

import (
	"Hybrid_Cloud/hcp-apiserver/pkg/util"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// gke container images
type Images struct {
	SRC_IMAGE  string
	DEST_IMAGE string
	IMAGE_NAME string
}

func (i *Images) AddTag(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "add-tag", i.SRC_IMAGE, i.DEST_IMAGE)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")
	if err != nil {
		log.Println(err)
	} else {
		w.Write(data)
	}
}

func (i *Images) Delete(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "delete", i.IMAGE_NAME)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) Describe(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "describe", i.IMAGE_NAME)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) List(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "list")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) ListTags(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "list-tags", i.IMAGE_NAME)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (i *Images) UnTags(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, i)
	cmd := exec.Command("gcloud", "container", "images", "untag", i.IMAGE_NAME)
	data, err := util.GetOutputReplaceStr(cmd, "Do you want to continue (Y/n)?", "")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

// operations
type Operations struct {
	OPERATION_ID string
}

func (o *Operations) Describe(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, o)
	cmd := exec.Command("gcloud", "container", "operations", "describe", o.OPERATION_ID)
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (o *Operations) List(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, o)
	cmd := exec.Command("gcloud", "container", "operations", "list")
	data, err := util.GetOutput(cmd)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(data))
		w.Write(data)
	}
}

func (o *Operations) Wait(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, o)
	cmd := exec.Command("gcloud", "container", "operations", "wait", o.OPERATION_ID)
	data, err := util.GetOutput(cmd)
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

func (d *Docker) Docker(w http.ResponseWriter, req *http.Request) {
	util.Parser(w, req, d)
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

func main() {
	var i Images
	http.HandleFunc("/gke/container/images/addTag", i.AddTag)
	http.HandleFunc("/gke/container/images/delete", i.Delete)
	http.HandleFunc("/gke/container/images/describe", i.Describe)
	http.HandleFunc("/gke/container/images/list", i.List)
	http.HandleFunc("/gke/container/images/listTags", i.ListTags)
	http.HandleFunc("/gke/container/images/unTags", i.UnTags)

	var o Operations
	http.HandleFunc("/gke/container/operations/describe", o.Describe)
	http.HandleFunc("/gke/container/operations/list", o.List)
	http.HandleFunc("/gke/container/operations/wait", o.Wait)

	var d Docker
	http.HandleFunc("/gke/docker", d.Docker)

	http.ListenAndServe(":3001", nil)
}
