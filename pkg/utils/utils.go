package utils

import (
	"errors"
	"fmt"
	"os"

	json "github.com/json-iterator/go"
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

// Marshal is a fast struct serializing method.
// Compared with encoding/json, it is claimed to have a better performance.
func Marshal(i interface{}) (string, error) {
	stream := json.ConfigFastest.BorrowStream(nil)
	defer json.ConfigFastest.ReturnStream(stream)

	stream.WriteVal(i)
	if stream.Error != nil {
		return "", errors.New("failed to marshal object")
	}
	return string(stream.Buffer()), nil
}

// Unmarshal is a fast struct deserializing method.
func Unmarshal(s string, obj interface{}) error {
	iter := json.ConfigFastest.BorrowIterator([]byte(s))
	defer json.ConfigFastest.ReturnIterator(iter)

	iter.ReadVal(obj)
	if iter.Error != nil {
		return errors.New("failed to unmarshal to object")
	}
	return nil
}
