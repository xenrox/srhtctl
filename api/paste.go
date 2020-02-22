package api

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

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

	return nil
}
