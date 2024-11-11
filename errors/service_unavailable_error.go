package errors

import (
	"errors"
)

type ServiceUnavailableError struct {
	message string
}

func (e *ServiceUnavailableError) Error() string {
	return e.message
}

func NewServiceUnavailableError(message string, args ...interface{}) Error {
	return &ServiceUnavailableError{message: formatMessage(message, args...)}
}

func IsServiceUnavailableError(err error) bool {
	var e *ServiceUnavailableError
	return errors.As(err, &e)
}
