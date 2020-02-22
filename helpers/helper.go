package helpers

import (
	"fmt"
)

// PrintError prints an error message
func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
