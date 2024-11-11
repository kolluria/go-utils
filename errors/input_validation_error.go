package errors

import (
	"errors"
)

type InputValidationError struct {
	message string
}

func (e *InputValidationError) Error() string {
	return e.message
}

func NewInputValidationError(message string, args ...interface{}) Error {
	return &InputValidationError{message: formatMessage(message, args...)}
}

func IsInputValidationError(err error) bool {
	var e *InputValidationError
	return errors.As(err, &e)
}
