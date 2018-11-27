package utils

import (
	"fmt"
	"os"
)

func PrintErr(msg string, err error) {
	fmt.Fprintf(os.Stderr, msg, err)
	os.Exit(1)
}
