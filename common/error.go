package common

import (
	"fmt"
	"strings"
)

// OperationErrorType - the error type
type OperationErrorType int8

// Error types
const (
	APIError OperationErrorType = iota
	NotFound
	ConfigError
	DatabaseError
	SystemError
	DataError
	ValidationError
	AvailabilityError
	ForbiddenError
	PaymentError
)

// Name - returns the name of the error type
func (oe OperationErrorType) Name() string {
	names := [...]string{
		"APIError",
		"NotFound",
		"ConfigError",
		"DatabaseError",
		"SystemError",
		"DataError",
		"ValidationError",
		"AvailabilityError",
		"ForbiddenError",
		"PaymentError",
	}

	// prevent panicking in case of
	// `name` is out of range
	if oe < APIError || oe > PaymentError {
		return "Unknown"
	}

	return names[oe]
}

// SE6367Error is a extension of error containing specific errors within API.
type SE6367Error struct {
	ErrorType OperationErrorType `json:"error_type"`
	Msg       string             `json:"error_msg"`
}

// NewAPIError - API error response wrapper
func NewAPIError(msg string) *SE6367Error {
	return &SE6367Error{APIError, fmt.Sprintf("API Error: %s", msg)}
}

// NewNotFoundError - API Not Found error response wrapper
func NewNotFoundError(msg string) *SE6367Error {
	return &SE6367Error{NotFound, msg}
}

// NewConfigurationError - Configuration error
func NewConfigurationError(msg string) *SE6367Error {
	return &SE6367Error{ConfigError, fmt.Sprintf("Configuration Error: %s", msg)}
}

// NewDatabaseOperationError returns a database error describing the error.
func NewDatabaseOperationError(msg string) *SE6367Error {
	return &SE6367Error{DatabaseError, fmt.Sprintf("Database Operation Error: %s", msg)}
}

// NewSystemError - returns a system error.
func NewSystemError(msg string) *SE6367Error {
	return &SE6367Error{SystemError, fmt.Sprintf("System Error: %s", msg)}
}

// NewValidationError - model validation error
func NewValidationError(msg string) *SE6367Error {
	return &SE6367Error{ValidationError, fmt.Sprintf("Validation Error: %s", msg)}
}

// NewValidationConcatError - model validation error with multiple errors
func NewValidationConcatError(errors []*SE6367Error) *SE6367Error {
	builder := strings.Builder{}
	for _, err := range errors {
		builder.WriteString(fmt.Sprintf("%s,", err.Error()))
	}

	return &SE6367Error{ValidationError, fmt.Sprintf("Validation Error:%s", strings.Replace(builder.String(), "Validation Error:", "", -1))}
}

// NewDataError - data error
func NewDataError(msg string) *SE6367Error {
	return &SE6367Error{DataError, fmt.Sprintf("Data Error: %s", msg)}
}

// NewAvailabilityError - data error
func NewAvailabilityError(msg string) *SE6367Error {
	return &SE6367Error{AvailabilityError, fmt.Sprintf("Availability Error: %s", msg)}
}

// NewForbiddenError - action is not permitted
func NewForbiddenError(msg string) *SE6367Error {
	return &SE6367Error{ForbiddenError, fmt.Sprintf("Forbidden: %s", msg)}
}

// NewPaymentError - action is not permitted
func NewPaymentError(msg string) *SE6367Error {
	return &SE6367Error{PaymentError, fmt.Sprintf("Payment Error: %s", msg)}
}

// NewSingletonErrorList - wraps error into a list
func NewSingletonErrorList(error *SE6367Error) []*SE6367Error {
	return []*SE6367Error{error}
}

// Implicit implement of Error interface
func (err *SE6367Error) Error() string {
	return err.Msg
}
