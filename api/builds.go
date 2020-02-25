package api

import (
	"errors"
	"fmt"
	"io/ioutil"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
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

type manifestResponseStruct struct {
	ID int `json:"id"`
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
		helpers.PrintError(err)
	}

	return nil
}

// BuildResubmit resubmits a build ID
func BuildResubmit(args []string) error {
	if len(args) != 1 {
		fmt.Println("Please append one build ID")
		return nil
	}
	var manifest string
	url := fmt.Sprintf("%s/api/jobs/%s/manifest", config.GetURL("builds"), args[0])
	err := Request(url, "GET", "", &manifest)
	if err != nil {
		return err
	}
	err = buildDeployManifest(manifest)
	if err != nil {
		return err
	}
	return nil
}

// BuildInformation gets information about a job by its ID
func BuildInformation(args []string) error {
	if len(args) != 1 {
		fmt.Println("Please append one build ID")
	}
	var response buildStruct
	url := fmt.Sprintf("%s/api/jobs/%s", config.GetURL("builds"), args[0])
	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	// TODO: write good print function that iterates through tasks
	fmt.Println(response)
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
