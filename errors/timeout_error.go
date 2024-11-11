package errors

import (
	"errors"
)

type TimeoutError struct {
	message string
}

func (e *TimeoutError) Error() string {
	return e.message
}

func NewTimeoutError(message string, args ...interface{}) Error {
	return &TimeoutError{message: formatMessage(message, args...)}
}

func IsTimeoutError(err error) bool {
	var e *TimeoutError
	return errors.As(err, &e)
}
