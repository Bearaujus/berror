package berror

import (
	"fmt"
	"strings"
)

// struct for error wrapper
type errorWrapper struct {
	format string
	code   string
}

// create new error wrapper instance
func NewErrorWrapper(format string, code ...string) *errorWrapper {
	if len(code) != 0 {
		return &errorWrapper{
			format: strings.Join([]string{"(%v)", format}, " "),
			code:   code[0],
		}
	}

	return &errorWrapper{
		format: format,
	}
}

// return new error from created error wrapper instance
func (ewi *errorWrapper) New(v ...interface{}) error {
	if ewi.code != "" {
		return fmt.Errorf(ewi.format, append([]interface{}{ewi.code}, v)...)
	}

	return fmt.Errorf(ewi.format, v...)
}
