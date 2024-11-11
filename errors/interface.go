package errors

import (
	"errors"
	"fmt"
)

// Error is an interface for representing execution errors.
// It is used to return errors from the API.
// Note: Any error returned by the internal methods should implement this interface.
type Error interface {
	error
}

func formatMessage(message string, args ...interface{}) string {
	if len(args) > 0 { // this check avoids unnecessary calls to fmt.Sprintf
		message = fmt.Sprintf(message, args...)
	}

	return message
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
