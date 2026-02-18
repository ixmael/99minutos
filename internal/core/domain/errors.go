package domain

import (
	"fmt"
	"time"
)

// ErrorCode represents a unique error identifier
type ErrorCode string

// ErrorType represents the category of error
type ErrorType string

const (
	ErrorTypeInfrastructure ErrorType = "INFRASTRUCTURE"
	ErrorTypeApplication    ErrorType = "APPLICATION"
	ErrorTypeDomain         ErrorType = "DOMAIN"
)

// BaseError is the common interface for all domain errors
type BaseError interface {
	error
	Code() ErrorCode
	Type() ErrorType
	Message() string
	Details() any
	Timestamp() time.Time
	Cause() error
	WithDetails(details any) BaseError
	WithCause(cause error) BaseError
	WithTimestamp(timestamp time.Time) BaseError
}

// BaseErrorImpl provides a default implementation for BaseError
type BaseErrorImpl struct {
	code      ErrorCode
	errorType ErrorType
	message   string
	details   any
	timestamp time.Time
	cause     error
}

func (e *BaseErrorImpl) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.code, e.message, e.cause)
	}
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

func (e *BaseErrorImpl) Code() ErrorCode {
	return e.code
}

func (e *BaseErrorImpl) Type() ErrorType {
	return e.errorType
}

func (e *BaseErrorImpl) Message() string {
	return e.message
}

func (e *BaseErrorImpl) Details() any {
	return e.details
}

func (e *BaseErrorImpl) Timestamp() time.Time {
	return e.timestamp
}

func (e *BaseErrorImpl) Cause() error {
	return e.cause
}

func (e *BaseErrorImpl) WithDetails(details any) BaseError {
	newError := *e
	newError.details = details
	return &newError
}

func (e *BaseErrorImpl) WithCause(cause error) BaseError {
	newError := *e
	newError.cause = cause
	return &newError
}

func (e *BaseErrorImpl) WithTimestamp(timestamp time.Time) BaseError {
	newError := *e
	newError.timestamp = timestamp
	return &newError
}

// NewBaseError creates a new base error with default timestamp
func NewBaseError(code ErrorCode, errorType ErrorType, message string) *BaseErrorImpl {
	return &BaseErrorImpl{
		code:      code,
		errorType: errorType,
		message:   message,
		timestamp: time.Now(),
	}
}
