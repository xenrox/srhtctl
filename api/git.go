package api

import (
	"errors"
	"fmt"
	"os/exec"

	"git.xenrox.net/~xenrox/srhtctl/config"
)

// GitUserName is the git user name without ~
var GitUserName string

// GitRepoName is the git repository name
var GitRepoName string

// GitAnnotate creates annotations
func GitAnnotate(args []string) error {
	if len(args) != 1 {
		return errors.New("Please append an annotations file")
	}
	cmd := exec.Command("git", "rev-parse", "HEAD")
	ref, err := cmd.Output()
	if err != nil {
		return errors.New("Cannot correctly execute git command. Not in a valid repository?")
	}
	// TODO: make repo and username required flags (cobra)
	url := fmt.Sprintf("%s/api/%s/repos/%s/%s/annotate", config.GetURL("git"), GitUserName, GitRepoName, ref)
	// FIXME: complete
	_ = url
	// err = Request(url, "PUT", "",)
	return nil
}
