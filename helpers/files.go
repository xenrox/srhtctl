package helpers

import (
	"errors"
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
