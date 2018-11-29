package utils

import (
	"fmt"
	"os"
)

// ErrExit is used to print error to stderr, and then exist with code 1.
func ErrExit(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg, err)
	} else {
		fmt.Fprintf(os.Stderr, msg)
	}
	os.Exit(1)
}
