package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/config"
)

type annotateResponseStruct struct {
	Updated int `json:"updated"`
}

// GitRepoName is the git repository name
var GitRepoName string

// GitAnnotate creates annotations
func GitAnnotate(args []string) error {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	ref, err := cmd.Output()
	if err != nil {
		return errors.New("Cannot correctly execute git command. Not in a valid repository?")
	}
	if config.UserName == "" {
		config.GetConfigValue("settings", "user")
	}
	url := fmt.Sprintf("%s/api/~%s/repos/%s/%s/annotate", config.GetURL("git"), config.UserName, GitRepoName, strings.TrimRight(string(ref), "\n"))
	var response annotateResponseStruct
	annotations, err := ioutil.ReadFile(args[0])
	err = FormRequest(url, "PUT", string(annotations), &response)
	if err != nil {
		return err
	}
	fmt.Printf("Added %d annotations.\n", response.Updated)
	return nil
}
