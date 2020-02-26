package api

import (
	"errors"
	"fmt"
	"io/ioutil"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
)

type buildDeployStruct struct {
	Manifest string `json:"manifest"`
	Note     string `json:"Note"`
	// Tags     []string `json:"tags"`
}

type taskStruct struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Log    string `json:"log"`
}

type buildStruct struct {
	ID       int          `json:"id"`
	Status   string       `json:"status"`
	SetupLog string       `json:"setup_log"`
	Tasks    []taskStruct `json:"tasks"`
	Note     string       `json:"note"`
	Tags     []string     `json:"tags"`
	Runner   string       `json:"runner"`
	Owner    struct {
		CName string `json:"canonical_name"`
		Name  string `json:"name"`
	} `json:"owner"`
}

// BuildDeploy deploys build manifests from command line
func BuildDeploy(args []string) error {
	if len(args) == 0 {
		return errors.New("Please append build manifests")
	}
	for _, file := range args {
		manifest, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = buildDeployManifest(string(manifest))
		errorhelper.PrintError(err)
	}

	return nil
}

// BuildResubmit resubmits a build ID
func BuildResubmit(args []string) error {
	if len(args) != 1 {
		return errors.New("Please append one build ID")
	}
	var manifest string
	url := fmt.Sprintf("%s/api/jobs/%s/manifest", config.GetURL("builds"), args[0])
	err := Request(url, "GET", "", &manifest)
	if err != nil {
		return err
	}
	return buildDeployManifest(manifest)
}

// BuildInformation gets information about a job by its ID
func BuildInformation(args []string) error {
	if len(args) != 1 {
		return errors.New("Please append one build ID")
	}
	var response buildStruct
	url := fmt.Sprintf("%s/api/jobs/%s", config.GetURL("builds"), args[0])
	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	printBuildInformation(response)
	return nil
}

func buildDeployManifest(manifest string) error {
	url := fmt.Sprintf("%s/api/jobs", config.GetURL("builds"))
	var body buildDeployStruct
	// TODO: parse tags and notes too with flags
	body.Manifest = manifest

	var response buildStruct
	err := Request(url, "POST", body, &response)
	if err != nil {
		return err
	}
	HandleResponse(fmt.Sprintf("%s/%s/job/%d\n", config.GetURL("builds"), response.Owner.CName, response.ID), true)
	return nil
}

func printBuildInformation(information buildStruct) {
	fmt.Printf("Build %d: %s\n", information.ID, formatBuildStatus(information.Status))
	for _, task := range information.Tasks {
		fmt.Printf("\tTask %s: %s\n", task.Name, formatBuildStatus(task.Status))
	}
	// TODO: if failed, and flag/config option set: append last lines of failed log
}

func formatBuildStatus(status string) string {
	switch status {
	case "running":
		status = fmt.Sprintf("\033[94m%s\033[0m", status)
	case "success":
		status = fmt.Sprintf("\033[92m%sful\033[0m", status)
	case "failed":
		status = fmt.Sprintf("\033[91m%s\033[0m", status)
	case "cancelled":
		status = fmt.Sprintf("\033[93m%s\033[0m", status)
	// default: pending or queued
	default:
		status = fmt.Sprintf("\033[1m%s\033[0m", status)
	}
	return status
}
