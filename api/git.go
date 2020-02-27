package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/config"
)

// GitUserName is the git user name without ~
var GitUserName string

// GitRepoName is the git repository name
var GitRepoName string

// GitAnnotate creates annotations
func GitAnnotate(args []string) error {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	ref, err := cmd.Output()
	if err != nil {
		return errors.New("Cannot correctly execute git command. Not in a valid repository?")
	}
	// TODO: make repo and username required flags (cobra)
	url := fmt.Sprintf("%s/api/~%s/repos/%s/%s/annotate", config.GetURL("git"), GitUserName, GitRepoName, strings.TrimRight(string(ref), "\n"))
	// TODO: correctly use json and print response in a sane way
	var response string
	annotations, err := ioutil.ReadFile(args[0])
	err = FormRequest(url, "PUT", string(annotations), &response)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
