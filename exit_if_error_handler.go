package berror

import (
	"fmt"
	"os"
)

// struct for exit if error handler
type exitIfErrorHandler struct {
	funcsIfErr []func() error
}

// create new exit if error handler
func NewexitIfErrorHandler(funcsIfErr ...func() error) *exitIfErrorHandler {
	return &exitIfErrorHandler{
		funcsIfErr: funcsIfErr,
	}
}

// if this function receive not nil error, then do os.Exit()
func (eh *exitIfErrorHandler) ExitIfError(err error, exitCode ...int) {
	if !IsError(err) {
		return
	}

	fmt.Println(err)
	if eh.hasFuncsIFErr() {
		if errFuncs := eh.executeFuncsIfErr(); errFuncs != nil {
			fmt.Println(errFuncs)
		}
	}
	
	if len(exitCode) != 0 {
		os.Exit(exitCode[0])
	}

	os.Exit(1)
}

// execute functions from instance funcs if err
func (eh *exitIfErrorHandler) executeFuncsIfErr() error {
	for _, v := range eh.funcsIfErr {
		if err := v(); err != nil {
			return err
		}
	}

	return nil
}

// check if funcs if err is not empty
func (eh *exitIfErrorHandler) hasFuncsIFErr() bool {
	return len(eh.funcsIfErr) != 0
}
