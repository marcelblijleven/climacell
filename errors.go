package climacell

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidBaseURL   = errors.New("invalid base URL provided")
	ErrInvalidAPIKey    = errors.New("invalid api key provided")
	ErrInvalidLatitude  = errors.New("invalid latitude provided")
	ErrInvalidLongitude = errors.New("invalid longitude provided")
)

// HTTPError represents an error that was returned from the climacell API
type HTTPError struct {
	Endpoint string `json:"-"`
	Message  string `json:"message"`
}

// Error returns the string representation of the HTTPError
func (e *HTTPError) Error() string {
	return fmt.Sprintf("error %v: %v", e.Endpoint, e.Message)
}

// BadRequestError represents an HTTP error with status code 400 - bad request
type BadRequestError struct {
	HTTPError
}

// Error returns the string representation of the BadRequestError
func (e *BadRequestError) Error() string {
	return fmt.Sprintf("bad request %v: %v", e.Endpoint, e.Message)
}

// UnauthorizedError represents an HTTP error with status code 401 - unauthorized
type UnauthorizedError struct {
	HTTPError
}

// Error returns the string representation of the UnauthorizedError
func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized %v: %v", e.Endpoint, e.Message)
}

// ForbiddenError represents an HTTP error with status code 403 - forbidden
type ForbiddenError struct {
	HTTPError
}

// Error returns the string representation of the ForbiddenError
func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("forbidden %v: %v", e.Endpoint, e.Message)
}

// TooManyRequestsError represents an HTTP error with status code 429 - too many requests
type TooManyRequestsError struct {
	HTTPError
}

// Error returns the string representation of the TooManyRequestsError
func (e *TooManyRequestsError) Error() string {
	return fmt.Sprintf("too many requests %v: %v", e.Endpoint, e.Message)
}

// InternalServerError represents an HTTP error with status code 500 - internal server error
type InternalServerError struct {
	HTTPError
}

// Error returns the string representation of the InternalServerError
func (e *InternalServerError) Error() string {
	return fmt.Sprintf("internal server error %v: %v", e.Endpoint, e.Message)
}

// NotFoundError represents an HTTP error with status code 404 - not found
type NotFoundError struct {
	HTTPError
}

// Error returns the string representation of the NotFoundError
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found error %v: %v", e.Endpoint, e.Message)
}

func newBadRequestError(endpoint, message string) *BadRequestError {
	return &BadRequestError{HTTPError{
		Endpoint: endpoint,
		Message:  message,
	}}
}

func newUnauthorizedError(endpoint, message string) *UnauthorizedError {
	return &UnauthorizedError{HTTPError{
		Endpoint: endpoint,
		Message:  message,
	}}
}

func newForbiddenError(endpoint, message string) *ForbiddenError {
	return &ForbiddenError{HTTPError{
		Endpoint: endpoint,
		Message:  message,
	}}
}

func newTooManyRequestError(endpoint, message string) *TooManyRequestsError {
	return &TooManyRequestsError{HTTPError{
		Endpoint: endpoint,
		Message:  message,
	}}
}

func newInternalServerError(endpoint, message string) *InternalServerError {
	return &InternalServerError{HTTPError{
		Endpoint: endpoint,
		Message:  message,
	}}
}

func newNotFoundError(endpoint, message string) *NotFoundError {
	return &NotFoundError{HTTPError{
		Endpoint: endpoint,
		Message:  message,
	}}
}
