package errorhelper

import (
	"errors"
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
// Parameter: error
func ExitError(err error) {
	if err != nil {
		PrintError(err)
		os.Exit(1)
	}
}

// ExitString calls PrintError and exits the program with error code
// Parameter: string
func ExitString(err string) {
	if err != "" {
		PrintError(errors.New(err))
		os.Exit(1)
	}
}
