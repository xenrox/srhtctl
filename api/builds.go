package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
	"git.xenrox.net/~xenrox/srhtctl/helpers/errorhelper"
)

type buildDeployStruct struct {
	Manifest string   `json:"manifest"`
	Note     string   `json:"note"`
	Tags     []string `json:"tags"`
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
	Note     *string      `json:"note"`
	Tags     *string      `json:"tags"`
	Runner   *string      `json:"runner"`
	Owner    struct {
		CName string `json:"canonical_name"`
		Name  string `json:"name"`
	} `json:"owner"`
}

// BuildNote is a description of a build
var BuildNote string

// BuildTags are tags for a build
var BuildTags []string

// BuildDeploy deploys build manifests from command line
func BuildDeploy(args []string) error {
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

// BuildEdit If true, then edit the manifest before resubmit
var BuildEdit bool

// BuildResubmit resubmits a build ID
func BuildResubmit(args []string) error {
	var buildDeploy buildDeployStruct
	var buildInfo buildStruct
	url := fmt.Sprintf("%s/api/jobs/%s/manifest", config.GetURL("builds"), args[0])
	err := Request(url, "GET", "", &buildDeploy.Manifest)
	if err != nil {
		return err
	}

	err = buildGetStruct(args[0], &buildInfo)
	if err != nil {
		return err
	}

	BuildNote = fmt.Sprintf("Resubmission of build [#%s](/~xenrox/job/%s)", args[0], args[0])
	BuildTags = helpers.TransformTags(*buildInfo.Tags)

	if BuildEdit {
		file, err := ioutil.TempFile(os.TempDir(), "srhtctl*.yml")
		if err != nil {
			return err
		}
		fileName := file.Name()
		defer os.Remove(fileName)
		_, err = file.WriteString(fmt.Sprintf("# Set new build note here: %s\n", BuildNote))
		if err != nil {
			return err
		}
		_, err = file.WriteString(buildDeploy.Manifest)
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
		err = helpers.EditFile(fileName)
		if err != nil {
			return err
		}
		fileContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}
		lines := strings.Split(string(fileContent), "\n")
		BuildNote = strings.Split(lines[0], ":")[1]
		return buildDeployManifest(string(fileContent))
	}
	return buildDeployManifest(buildDeploy.Manifest)
}

// BuildInformation gets information about a job by its ID
func BuildInformation(args []string) error {
	var response buildStruct
	err := buildGetStruct(args[0], &response)
	if err != nil {
		return err
	}
	return response.printBuildInformation()
}

func buildGetStruct(number string, response *buildStruct) error {
	url := fmt.Sprintf("%s/api/jobs/%s", config.GetURL("builds"), number)
	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	return nil
}

func buildDeployManifest(manifest string) error {
	url := fmt.Sprintf("%s/api/jobs", config.GetURL("builds"))
	var body buildDeployStruct
	body.Manifest = manifest
	body.Note = BuildNote
	if len(BuildTags) > 0 {
		body.Tags = BuildTags
	} else {
		body.Tags = make([]string, 0)
	}
	var response buildStruct
	err := Request(url, "POST", body, &response)
	if err != nil {
		return err
	}
	HandleResponse(fmt.Sprintf("%s/%s/job/%d\n", config.GetURL("builds"), response.Owner.CName, response.ID), true)
	return nil
}

func (information buildStruct) printBuildInformation() error {
	fmt.Printf("Build %d: %s\n", information.ID, formatBuildStatus(information.Status))
	for _, task := range information.Tasks {
		fmt.Printf("\tTask %s: %s\n", task.Name, formatBuildStatus(task.Status))
	}
	if information.Status == "failed" {
		debugLines := config.GetConfigValue("builds", "debugLines", "0")
		if debugLines != "0" {
			length, err := strconv.Atoi(debugLines)
			if err != nil {
				return err
			}
			fmt.Printf("\n\n\033[4mBuild setup failed with:\033[0m\n\n")
			err = printBuildErrors(information.SetupLog, length)
			if err != nil {
				return err
			}
			for _, task := range information.Tasks {
				if task.Status == "failed" {
					fmt.Printf("\n\033[4mTask %s failed with:\033[0m\n\n", task.Name)
					err = printBuildErrors(task.Log, length)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func printBuildErrors(url string, length int) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	lines := strings.Split(string(responseBody), "\n")
	for _, line := range lines[helpers.Max(len(lines)-length, 0):] {
		fmt.Println(line)
	}
	return nil
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
