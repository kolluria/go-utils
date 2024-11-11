package errors

import (
	"errors"
)

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

func NewNotFoundError(message string, args ...interface{}) Error {
	return &NotFoundError{message: formatMessage(message, args...)}
}

func IsNotFoundError(err error) bool {
	var e *NotFoundError
	return errors.As(err, &e)
}
