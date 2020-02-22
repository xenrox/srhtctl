package helpers

import (
	"encoding/json"
	"fmt"
)

// PrintError prints an error message
func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// EscapeJSON formats a string so that it can be used in a json struct
func EscapeJSON(content string) (string, error) {
	escaped, err := json.Marshal(content)
	if err != nil {
		return "", err
	}
	return string(escaped[1 : len(escaped)-1]), nil
}
