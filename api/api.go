package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"github.com/atotto/clipboard"
)

// Request does the actual API request
func Request(url string, method string, body interface{}, response interface{}) error {
	client := &http.Client{Timeout: 3 * time.Second}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	token := config.GetConfigValue("settings", "token")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(responseBody, &response)
	return nil
}

// HandleResponse prints the response and can copy it to clipboard
func HandleResponse(response string) {
	if response == "" {
		return
	}
	if config.GetConfigValue("settings", "copyToClipboard", "false") == "true" {
		clipboard.WriteAll(response)
	}
	fmt.Println(response)
}
