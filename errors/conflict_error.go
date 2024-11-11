package errors

import (
	"errors"
)

type ConflictError struct {
	message string
}

func (e *ConflictError) Error() string {
	return e.message
}

func NewConflictError(message string, args ...interface{}) Error {
	return &ConflictError{message: formatMessage(message, args...)}
}

func IsConflictError(err error) bool {
	var conflictError *ConflictError
	return errors.As(err, &conflictError)
}
