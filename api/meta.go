package api

import (
	"fmt"
	"time"

	"git.xenrox.net/~xenrox/srhtctl/config"
)

type metaUserStruct struct {
	userStruct
	UsePGPKey *int `json:"use_pgp_key"`
}

type metaAudit struct {
	ID      int       `json:"id"`
	IP      string    `json:"ip"`
	Action  string    `json:"action"`
	Details string    `json:"details"`
	Created time.Time `json:"created"`
}

type metaAuditPagination struct {
	pagination
	Results []metaAudit `json:"results"`
}

type sshKeyStruct struct {
	ID          int             `json:"id"`
	Authorized  time.Time       `json:"authorized"`
	Comment     string          `json:"comment"`
	Fingerprint string          `json:"fingerprint"`
	Key         string          `json:"key"`
	Owner       shortUserStruct `json:"owner"`
	LastUsed    time.Time       `json:"last_used"`
}

type pgpKeyStruct struct {
	ID         int             `json:"id"`
	Key        string          `json:"key"`
	KeyID      string          `json:"key_id"`
	Email      string          `json:"email"`
	Authorized time.Time       `json:"authorized"`
	Owner      shortUserStruct `json:"owner"`
}

// MetaEdit If true, then edit own profile information
var MetaEdit bool

// MetaGetProfile prints profile information
func MetaGetProfile() error {
	url := fmt.Sprintf("%s/api/user/profile", config.GetURL("meta"))
	var response metaUserStruct
	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	response.printProfile()
	if MetaEdit == true {
		// TODO: edit values in text file vs with command line options
	}
	return nil
}

// MetaGetLogs prints audit log entries
func MetaGetLogs() error {
	// TODO allow user to iterate through pagination
	url := fmt.Sprintf("%s/api/user/audit-log", config.GetURL("meta"))
	var response metaAuditPagination
	err := Request(url, "GET", "", &response)
	if err != nil {
		return err
	}
	// reverse order: new entries at bottom of terminal
	for i := len(response.Results) - 1; i >= 0; i-- {
		response.Results[i].printAudit()
	}
	return nil
}

func (information metaUserStruct) printProfile() {
	fmt.Printf("Username: %s\n", information.Name)
	fmt.Printf("Email: %s\n", information.Email)
	if information.URL != nil {
		fmt.Printf("Url: %s\n", *information.URL)
	}
	if information.Location != nil {
		fmt.Printf("Location: %s\n", *information.Location)
	}
	if information.Bio != nil {
		fmt.Printf("Bio: %s\n", *information.Bio)
	}
	if information.UsePGPKey != nil {
		fmt.Printf("UsePGPKey: %d\n", *information.UsePGPKey)
	}
}

func (information metaAudit) printAudit() {
	fmt.Printf("In audit log %d (%s) the IP %s ", information.ID, information.Created.Format("2006/01/02 15:04:05"), information.IP)
	fmt.Printf("%s: %s\n", information.Action, information.Details)
}
