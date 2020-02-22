package api

import (
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
		fmt.Println("Please append build manifests")
	} else {
		for _, file := range args {
			err := buildDeployManifest(file)
			helpers.PrintError(err)
		}
	}
	return nil
}

func buildDeployManifest(file string) error {
	url := fmt.Sprintf("%s/api/jobs", config.GetURL("builds"))
	manifestContent, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var body manifestStruct
	// TODO: parse tags and notes too with flags
	body.Manifest = string(manifestContent)

	var response manifestResponseStruct
	err = Request(url, "POST", body, &response)
	if err != nil {
		return err
	}
	HandleResponse(fmt.Sprintf("Deployed build with build ID %s to %s", strconv.Itoa(response.ID), config.GetURL("builds")), false)
	return nil
}
