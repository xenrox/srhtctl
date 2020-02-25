package helpers

import (
	"fmt"
	"os"
)

// PrintError prints an error message to stderr
func PrintError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

// ExitError calls PrintError and exits the program with error code
func ExitError(err error) {
	PrintError(err)
	os.Exit(1)
}
