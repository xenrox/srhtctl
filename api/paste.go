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

// PasteFile is the path to the file that should be uploaded
var PasteFile string

// PasteName is the name of the file that should be uploaded
var PasteName string

// PasteVisibility is the visibility of the file that should be uploaded
var PasteVisibility string

// PasteExpiration is the duration that the paste should live
var PasteExpiration string

// PasteCreate creates a new paste resource
func PasteCreate() error {
	var escapedContent, name string

	if PasteFile != "" {
		fileContent, err := ioutil.ReadFile(PasteFile)
		if err != nil {
			return err
		}
		escapedContent, err = helpers.EscapeJSON(string(fileContent))
		if err != nil {
			return err
		}
		name = filepath.Base(PasteFile)
	} else {
		// paste from clipboard
		stringContent, err := clipboard.ReadAll()
		if err != nil {
			return err
		}
		escapedContent, err = helpers.EscapeJSON(stringContent)
		name = PasteName
	}

	var visibility string
	// TODO: check for illegal values
	if PasteVisibility != "" {
		visibility = PasteVisibility
	} else {
		visibility = config.GetConfigValue("paste", "visibility", "unlisted")
	}

	url := fmt.Sprintf("%s/api/pastes", config.GetURL("paste"))

	var body pasteCreateStruct
	body.Visibility = visibility
	body.Files = make([]pasteFileStruct, 1)
	body.Files[0].Filename = name
	body.Files[0].Contents = escapedContent

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
