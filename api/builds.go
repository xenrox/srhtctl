package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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
	Note     string       `json:"note"`
	Tags     *string      `json:"tags"`
	Runner   *string      `json:"runner"`
	Owner    userStruct   `json:"owner"`
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

	BuildNote = fmt.Sprintf("Resubmission of build [#%s](/%s/job/%s)", args[0], buildInfo.Owner.CName, args[0])
	BuildTags = helpers.TransformTags(buildInfo.Tags)

	if BuildEdit {
		file, fileName, err := helpers.CreateFile("srhtctl*.yml")
		if err != nil {
			return err
		}
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
		BuildNote = strings.SplitN(lines[0], ":", 2)[1]
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

	fmt.Println(response)
	return nil
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

func (task taskStruct) String() string {
	str := fmt.Sprintf("Task %s: %s\n", task.Name, formatBuildStatus(task.Status))

	return str
}

func (build buildStruct) String() string {
	str := fmt.Sprintf("Build %d: %s\n", build.ID, formatBuildStatus(build.Status))
	for _, task := range build.Tasks {
		str += fmt.Sprintf("\t%s", task)
	}

	if build.Status == "running" {
		log, err := retrieveBuildLog(build.SetupLog)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not retrieve logs")
			return str
		}

		sshCommand := getSSHCommand(log)

		if sshCommand != "" {
			str = str + "\n" + sshCommand
		}
	}

	if build.Status == "failed" {
		debugLines := config.GetConfigValue("builds", "debugLines", "0")

		setupLog, err := retrieveBuildLog(build.SetupLog)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not retrieve logs")
			return str
		}

		if debugLines != "0" {
			length, err := strconv.Atoi(debugLines)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Config: debugLines is no valid integer")
				return str
			}

			str += fmt.Sprintf("\n\n\033[4mBuild setup failed with:\033[0m\n\n")
			str += getBuildErrors(setupLog, length)
			for _, task := range build.Tasks {
				if task.Status == "failed" {
					str += fmt.Sprintf("\n\033[4mTask %s failed with:\033[0m\n\n", task.Name)
					taskLog, err := retrieveBuildLog(task.Log)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Could no retrieve logs")
						return str
					}

					str += getBuildErrors(taskLog, length)
				}
			}
		}

		str += getSSHCommand(setupLog)
	}

	return str
}

// getSSHCommand searches for the SSH connection command in log and returns it
func getSSHCommand(log string) string {
	var sshRegex = regexp.MustCompile("ssh -t .* [0-9]*")
	var str string

	sshCommand := sshRegex.FindString(log)
	if sshCommand != "" {
		str = fmt.Sprintf("SSH login command: \033[96m%s\033[0m", sshCommand)
		str += helpers.CopyToClipboard(sshCommand)
	}
	return str
}

// retrieveBuildLog gets the log file for a build task und returns it as a string
func retrieveBuildLog(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func getBuildErrors(log string, length int) string {
	var str string

	lines := strings.Split(log, "\n")
	for _, line := range lines[helpers.Max(len(lines)-length, 0):] {
		str += fmt.Sprintln(line)
	}

	return str
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
