package errors

import (
	"errors"
)

type NotImplementedError struct {
	message string
}

func (e *NotImplementedError) Error() string {
	return e.message
}

func NewNotImplementedError(message string, args ...interface{}) Error {
	return &NotImplementedError{message: formatMessage(message, args...)}
}

func IsNotImplementedError(err error) bool {
	var e *NotImplementedError
	return errors.As(err, &e)
}
