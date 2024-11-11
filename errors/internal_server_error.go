package errors

import (
	"errors"
)

type InternalServerError struct {
	message string
}

func (e *InternalServerError) Error() string {
	return e.message
}

func NewInternalServerError(message string, args ...interface{}) Error {
	return &InternalServerError{message: formatMessage(message, args...)}
}

func IsInternalServerError(err error) bool {
	var e *InternalServerError
	return errors.As(err, &e)
}
