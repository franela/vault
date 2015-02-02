package ui

import (
	"fmt"
	"os"
)

var (
	DEBUG = false
)

var message string

func Printf(format string, a ...interface{}) (n int, err error) {

	if DEBUG {
		newMessage := fmt.Sprintf(format, a...)
		message = message + newMessage
		return len(newMessage), nil
	} else {
		return fmt.Printf(format, a...)
	}

}

func PrintErrorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

func GetOutput() string {
	defer func() {
		message = ""
	}()

	return message
}
