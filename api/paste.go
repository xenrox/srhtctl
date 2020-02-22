package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
	"github.com/atotto/clipboard"
)

type pasteStruct struct {
	Created    string `json:"created"`
	Visibility string `json:"visibilty"`
	SHA        string `json:"sha"`
	User       struct {
		CName string `json:"canonical_name"`
		Name  string `json:"name"`
	} `json:"user"`
	Files []struct {
		Filename string `json:"filename"`
		BlobID   string `json:"blob_id"`
	} `json:"files"`
}

type pasteCreateStruct struct {
	Visibility string            `json:"visibility"`
	Files      []pasteFileStruct `json:"files"`
}

type pasteFileStruct struct {
	Contents string `json:"contents"`
	Filename string `json:"filename"`
}

// PasteName is the name of the file that should be uploaded
var PasteName string

// PasteVisibility is the visibility of the file that should be uploaded
var PasteVisibility string

// PasteExpiration is the duration that the paste should live
var PasteExpiration string

// PasteCreate creates a new paste resource
func PasteCreate(args []string) error {
	var body pasteCreateStruct

	// TODO: refactor

	if len(args) > 0 {
		body.Files = make([]pasteFileStruct, len(args))

		for i, file := range args {
			body.Files[i].Filename = filepath.Base(file)
			fileContent, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}
			body.Files[i].Contents = string(fileContent)
		}
	} else {
		// paste from clipboard
		stringContent, err := clipboard.ReadAll()
		if err != nil {
			return err
		}

		body.Files = make([]pasteFileStruct, 1)
		body.Files[0].Filename = PasteName
		body.Files[0].Contents = stringContent
	}

	var visibility string
	// TODO: check for illegal values
	if PasteVisibility != "" {
		visibility = PasteVisibility
	} else {
		visibility = config.GetConfigValue("paste", "visibility", "unlisted")
	}

	url := fmt.Sprintf("%s/api/pastes", config.GetURL("paste"))

	body.Visibility = visibility

	var response pasteStruct
	err := Request(url, "POST", body, &response)
	if err != nil {
		return err
	}

	var expiration string
	if PasteExpiration != "" {
		expiration = PasteExpiration
	} else {
		expiration = config.GetConfigValue("paste", "expiration", "0")
	}
	if expiration != "0" {
		logPath, err := pasteGetLog()
		helpers.PrintError(err)
		f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		helpers.PrintError(err)
		defer f.Close()

		sec := int(time.Now().Unix())
		expirationDays, err := strconv.Atoi(expiration)
		helpers.PrintError(err)
		sec += expirationDays * 24 * 60 * 60
		timeStamp := strconv.Itoa(sec)
		_, err = f.WriteString(fmt.Sprintf("%s:%s\n", response.SHA, timeStamp))
		helpers.PrintError(err)
	}

	HandleResponse(fmt.Sprintf("%s/%s/%s\n", config.GetURL("paste"), response.User.CName, response.SHA))

	return nil
}

// PasteDelete deletes multiple paste resources
func PasteDelete(args []string) error {
	// TODO: maybe delete from clipboard too
	for _, sha := range args {
		err := pasteDeleteSHA(sha)
		helpers.PrintError(err)
	}
	return nil
}

func pasteDeleteSHA(sha string) error {
	// TODO: feedback if paste was already deleted
	url := fmt.Sprintf("%s/api/pastes/%s", config.GetURL("paste"), sha)
	body := ""
	err := Request(url, "DELETE", body)
	if err != nil {
		return err
	}
	return nil
}

func pasteGetLog() (string, error) {
	logPath := config.GetConfigValue("paste", "logfile", "")
	if logPath == "" {
		xdgConfigHome, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		logPath = fmt.Sprintf("%s/srhtctl/pastes.log", xdgConfigHome)
	}
	return logPath, nil
}
