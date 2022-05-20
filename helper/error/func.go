package errConv

import (
	"errors"
	"strings"
)

func Conversion(err error) error {
	var errNew error

	if strings.Contains(err.Error(), "not found") {
		errNew = errors.New(ErrDBNotFound)
	}

	if strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "missmatch") || strings.Contains(err.Error(), "failure") {
		errNew = errors.New(ErrUserFailure)
	}

	return errNew
}
