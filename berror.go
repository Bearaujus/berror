package berror

import (
	"fmt"
	"os"
)

// Check if error is not nil
func IsError(err error) bool {
	return err != nil
}

// If this function receive not nil error, than print that error
func IsErrorPrint(err error) bool {
	isError := IsError(err)
	if isError {
		fmt.Println(err)
	}
	return isError
}

// If this function receive not nil error, than do os.exit()
func ExitIfError(err error, exitCode ...int) {
	if IsError(err) {
		fmt.Println(err)
		if len(exitCode) != 0 {
			os.Exit(exitCode[0])
		}
		os.Exit(1)
	}
}
