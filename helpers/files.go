package helpers

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"

	"git.xenrox.net/~xenrox/srhtctl/config"
)

// EditFile opens a file in an editor
func EditFile(fileName string) error {
	editor := config.GetConfigValue("settings", "editor", "")
	if editor == "" {
		editor = os.Getenv("EDITOR")
		if editor == "" {
			return errors.New("Please set up an editor in your config.ini")
		}
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		return err
	}
	cmd := exec.Command(editorPath, fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// CreateFile creates a temporary file with the name bases on a glob
func CreateFile(glob string) (*os.File, string, error) {
	file, err := ioutil.TempFile(os.TempDir(), glob)
	if err != nil {
		return nil, "", err
	}

	return file, file.Name(), nil
}
