package helpers

import (
	"fmt"
	"os"
)

// PrintError prints an error message to stderr
func PrintError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
