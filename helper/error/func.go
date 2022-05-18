package errConv

import (
	"errors"
	"strings"
)

func Conversion(err error) error {
	var errNew error
	// var errCode int
	// var errObj map[string]interface{}
	if strings.Contains(err.Error(), "not found") {
		errNew = errors.New(ErrDBNotFound)
		// errCode = 404
	}
	if strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "missmatch") || strings.Contains(err.Error(), "failure") {
		errNew = errors.New(ErrUserFailure)
		// errCode = 402
	}

	// errObj["message"] = errNew
	// errObj["resCode"] = errCode

	return errNew
}
