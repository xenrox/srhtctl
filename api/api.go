package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"git.xenrox.net/~xenrox/srhtctl/helpers"
)

type userStruct struct {
	CName    string  `json:"canonical_name"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	URL      *string `json:"url"`
	Location *string `json:"location"`
	Bio      *string `json:"bio"`
}

type apiErrorResponse struct {
	Errors []apiError `json:"errors"`
}

type apiError struct {
	Reason string `json:"reason"`
	Field  string `json:"field"`
}

type pagination struct {
	Next           *int `json:"next"`
	ResultsPerPage int  `json:"results_per_page"`
	Total          int  `json:"total"`
}

func (e apiErrorResponse) Error() string {
	err := "Error respose from the API:"
	for _, errors := range e.Errors {
		if errors.Field != "" {
			err += fmt.Sprintf("\n%s: %s", errors.Field, errors.Reason)

		} else {
			err += fmt.Sprintf("\n%s", errors.Reason)
		}
	}
	return err
}

// Request does the actual API request
func Request(url string, method string, body interface{}, response ...interface{}) error {
	client := &http.Client{Timeout: 3 * time.Second}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	token := config.GetToken()
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

	if resp.StatusCode >= 400 {
		var errorResponse apiErrorResponse
		json.Unmarshal(responseBody, &errorResponse)
		return errorResponse
	}

	if len(response) > 0 {
		switch v := response[0].(type) {
		case *string:
			*v = string(responseBody)
		default:
			json.Unmarshal(responseBody, v)
		}
	}
	return nil
}

// FormRequest does an API request with x-www-form-urlencoded
func FormRequest(url string, method string, annotations string, response ...interface{}) error {
	client := &http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest(method, url, strings.NewReader(annotations))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "x-www-form-urlencoded")
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

	if resp.StatusCode >= 400 {
		var errorResponse apiErrorResponse
		json.Unmarshal(responseBody, &errorResponse)
		return errorResponse
	}

	if len(response) > 0 {
		switch v := response[0].(type) {
		case *string:
			*v = string(responseBody)
		default:
			json.Unmarshal(responseBody, v)
		}
	}
	return nil
}

// HandleResponse prints the response and can copy it to clipboard
func HandleResponse(response string, copy bool) {
	if response == "" {
		return
	}
	if copy {
		helpers.CopyToClipboard(response)
	}
	fmt.Println(response)
}
