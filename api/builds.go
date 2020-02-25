package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
)

type manifestStruct struct {
	Manifest string `json:"manifest"`
	Note     string `json:"Note"`
	// Tags     []string `json:"tags"`
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

func buildDeployManifest(manifest string) error {
	url := fmt.Sprintf("%s/api/jobs", config.GetURL("builds"))
	var body manifestStruct
	// TODO: parse tags and notes too with flags
	body.Manifest = manifest

	var response manifestResponseStruct
	err := Request(url, "POST", body, &response)
	if err != nil {
		return err
	}
	// TODO: send upstream patch for a better response with username
	HandleResponse(fmt.Sprintf("Deployed build with build ID %s to %s", strconv.Itoa(response.ID), config.GetURL("builds")), false)
	return nil
}
