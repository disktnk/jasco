package jasco

import (
	"fmt"
	"net/http"
)

const (
	requestURLNotFoundErrorCode = "J0001"
	internalServerErrorCode     = "J0002"
	requestBodyReadErrorCode    = "J0003"
	requestBodyParseErrorCode   = "J0004"
)

const (
	// SomethingWentWrong is an error message that doesn't explain anything
	// to users except something went wrong.
	SomethingWentWrong = "Something went wrong. Please try again later."
)

// NewInternalServerError creates an internal server error response having
// an error information.
func NewInternalServerError(err error) *Error {
	return NewError(internalServerErrorCode, SomethingWentWrong,
		http.StatusInternalServerError, err)
}

// Error has the error information reported to clients of API.
type Error struct {
	// Code has an error code of this error. It always have
	// a single alphabet prefix and 4-digit error number.
	Code string `json:"code"`

	// Message contains a user-friendly error message for users
	// of WebUI or clients of API. It MUST NOT have internal
	// error information nor debug information which are
	// completely useless for them.
	Message string `json:"message"`

	// RequestID has an ID of the request which caused this error.
	// This value is convenient for users or clients to report their
	// problems to the development team. The team can look up incidents
	// easily by greping logs with the ID.
	//
	// The type of RequestID is string because it can be
	// very large and JavaScript might not be able to handle
	// it as an integer.
	RequestID string `json:"request_id"`

	// Status has an appropriate HTTP status code for this error.
	Status int `json:"-"`

	// Err is an internal error information which should not be shown to users
	// directly.
	Err error `json:"-"`

	// Meta contains arbitrary information of the error.
	// What the information means depends on each error.
	// For instance, a validation error might contain
	// error information of each field.
	Meta map[string]interface{} `json:"meta"`
}

// NewError creates a new Error instance.
func NewError(code string, msg string, status int, err error) *Error {
	return &Error{
		Code:    code,
		Message: msg,
		Status:  status,
		Err:     err,
		Meta:    map[string]interface{}{},
	}
}

// SetRequestID set the ID of the current request to Error.
func (e *Error) SetRequestID(id uint64) {
	e.RequestID = fmt.Sprint(id)
}
