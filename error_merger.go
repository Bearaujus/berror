package berror

import (
	"errors"
	"strings"
)

// struct for error merger
type errorMerger struct {
	data []string
}

// create new error merger instance
func NewErrorMerger() *errorMerger {
	return &errorMerger{}
}

// if this function receive not nil error, then add that error to the error merger instance
func (em *errorMerger) AddIfNotNil(err error) bool {
	if err != nil {
		em.data = append(em.data, err.Error())

		return true
	}

	return false
}

// return all errors from error merger instance (merged into one error)
func (em *errorMerger) GetError(errSep ...string) error {
	if len(errSep) != 0 {
		return errors.New(strings.Join(em.data, errSep[0]))
	}

	return errors.New(strings.Join(em.data, "\n"))
}

// return all errors from error merger instance
func (em *errorMerger) GetErrors() []error {
	var res []error
	for _, v := range em.data {
		res = append(res, errors.New(v))
	}

	return res
}

// check if error merger instance has an error
func (em *errorMerger) HasError() bool {
	return len(em.data) != 0
}
