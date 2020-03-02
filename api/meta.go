package api

import (
	"fmt"

	"git.xenrox.net/~xenrox/srhtctl/config"
)

type metaUserStruct struct {
	userStruct
	UsePGPKey *int `json:"use_pgp_key"`
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
