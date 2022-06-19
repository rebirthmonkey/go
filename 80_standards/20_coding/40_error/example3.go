package main

import (
	"fmt"

	"github.com/rebirthmonkey/go/80_standards/20_coding/40_error/errcode"
	"github.com/rebirthmonkey/pkg/errors"
)

func main() {
	if err := bindUser(); err != nil {
		// %s: Returns the user-safe error string mapped to the error code or the error message if none is specified.
		fmt.Println("====================> %s <====================")
		fmt.Printf("%s\n\n", err)

		// %v: Alias for %s.
		fmt.Println("====================> %v <====================")
		fmt.Printf("%v\n\n", err)

		// %-v: Output caller details, useful for troubleshooting.
		fmt.Println("====================> %-v <====================")
		fmt.Printf("%-v\n\n", err)

		// %+v: Output full error stack details, useful for debugging.
		fmt.Println("====================> %+v <====================")
		fmt.Printf("%+v\n\n", err)

		// %#-v: Output caller details, useful for troubleshooting with JSON formatted output.
		fmt.Println("====================> %#-v <====================")
		fmt.Printf("%#-v\n\n", err)

		// %#+v: Output full error stack details, useful for debugging with JSON formatted output.
		fmt.Println("====================> %#+v <====================")
		fmt.Printf("%#+v\n\n", err)

		if errors.IsCode(err, errcode.ErrDatabase) {
			fmt.Println("this is a ErrDatabase error")
		}

		if errors.IsCode(err, errcode.ErrEncodingFailed) {
			fmt.Println("this is a ErrEncodingFailed error")
		}

		fmt.Println("the cause error is:", errors.Cause(err))
	}
}

func bindUser() error {
	if err := getUser(); err != nil {
		// Step3: Wrap the error with a new error message and a new error code if needed.
		return errors.WrapC(err, errcode.ErrEncodingFailed, "encoding user failed.")
	}
	return nil
}

func getUser() error {
	if err := queryDatabase(); err != nil {
		// Step2: Wrap the error with a new error message.
		return errors.Wrap(err, "getUser() failed.")
	}
	return nil
}

func queryDatabase() error {
	// Step1. Create error with specified error code.
	return errors.WithCode(errcode.ErrDatabase, "user 'XXX' not found.")
}
