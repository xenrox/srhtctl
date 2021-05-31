package api

import (
	"fmt"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/config"
)

// ListName is the name of a mailing list
var ListName string

// PatchStatus is the status of a patchset
var PatchStatus string

// PatchPrefix is the prefix of a patchset
var PatchPrefix string

type emailStruct struct {
	ID        int          `json:"id"`
	Created   string       `json:"created"`
	Subject   string       `json:"subject"`
	MessageID string       `json:"message_id"`
	ParentID  *int         `json:"parent_id"`
	ThreadID  *int         `json:"thread_it"`
	List      listsStruct  `json:"list"`
	Sender    *userStruct  `json:"sender"`
	Patchset  *patchStruct `json:"patchset"`
}

type emailPagerStruct struct {
	Next           *int           `json:"next,string"`
	Results        []*emailStruct `json:"results"`
	Total          int            `json:"total"`
	ResultsPerPage int            `json:"results_per_page"`
}

type patchStruct struct {
	ID        int     `json:"id"`
	Created   string  `json:"created"`
	Updated   string  `json:"updated"`
	Subject   string  `json:"subject"`
	Prefix    string  `json:"prefix"`
	Version   int     `json:"version"`
	Status    string  `json:"status"`
	Submitter string  `json:"submitter"`
	ReplyTo   *string `json:"reply_to"`
	MessageID string  `json:"message_id"`
}

type patchPagerStruct struct {
	Next           *int           `json:"next,string"`
	Results        []*patchStruct `json:"results"`
	Total          int            `json:"total"`
	ResultsPerPage int            `json:"results_per_page"`
}

type listPermissions struct {
	NonSubscriber []string `json:"nonsubscriber"`
	Subscriber    []string `json:"subscriber"`
	Account       []string `json:"account"`
}

type listsStruct struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Owner       userStruct      `json:"owner"`
	Created     string          `json:"created"`
	Updated     string          `json:"updated"`
	Description string          `json:"description"`
	Permissions listPermissions `json:"permissions"`
}

type listsPagerStruct struct {
	Next           *int           `json:"next,string"`
	Results        []*listsStruct `json:"results"`
	Total          int            `json:"total"`
	ResultsPerPage int            `json:"results_per_page"`
}

// PrintPatchsets prints out patchsets
func PrintPatchsets(args []string) error {
	var patches patchPagerStruct

	if ListName != "" {
		err := getPatchsets(&patches, ListName)
		if err != nil {
			return err
		}

		for _, patch := range patches.Results {
			fmt.Print(patch.filterByPrefix().filterByStatus())
		}
		return nil
	}

	var lists listsPagerStruct
	err := getLists(&lists)
	if err != nil {
		return err
	}

	for _, list := range lists.Results {
		err = getPatchsets(&patches, list.Name)
		if err != nil {
			return err
		}

		for i, patch := range patches.Results {
			if i == 0 {
				fmt.Printf("List %s:\n\n", list.Name)
			}
			fmt.Print(patch.filterByPrefix().filterByStatus())
		}
	}

	return nil
}

func getPatchsets(response *patchPagerStruct, listName string) error {
	url := fmt.Sprintf("%s/api/lists/%s/patchsets", config.GetURL("lists"), listName)
	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	return nil
}

func getLists(response *listsPagerStruct) error {
	url := fmt.Sprintf("%s/api/lists", config.GetURL("lists"))
	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	return nil
}

func (patch patchStruct) String() string {
	if patch == (patchStruct{}) {
		return ""
	}

	prefix := patch.Prefix

	var version string
	if patch.Version != 1 {
		version = fmt.Sprintf("v%d", patch.Version)
	}

	var suffix string
	if prefix != "" || version != "" {
		suffix = fmt.Sprintf("%s %s", prefix, version)
		suffix = fmt.Sprintf(" [%s]", strings.TrimSpace(suffix))
	}

	str := fmt.Sprintf("Patchset %d%s: %s\n", patch.ID, suffix, patch.Subject)
	str += fmt.Sprintf("Submitter: %s\n", patch.Submitter)
	str += fmt.Sprintf("Status: %s\n", patch.Status)

	str += "\n"
	return str
}

func (patch patchStruct) filterByStatus() patchStruct {
	if PatchStatus == "all" || PatchStatus == patch.Status {
		return patch
	}
	return patchStruct{}
}

func (patch patchStruct) filterByPrefix() patchStruct {
	if PatchPrefix == "" || PatchPrefix == patch.Prefix {
		return patch
	}
	return patchStruct{}
}
