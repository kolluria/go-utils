package errors

import (
	"errors"
)

type MethodNotAllowedError struct {
	message string
}

func (e *MethodNotAllowedError) Error() string {
	return e.message
}

func NewMethodNotAllowedError(message string, args ...interface{}) Error {
	return &MethodNotAllowedError{message: formatMessage(message, args...)}
}

func IsMethodNotAllowedError(err error) bool {
	var methodNotAllowedError *MethodNotAllowedError
	return errors.As(err, &methodNotAllowedError)
}
