package mandrill

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lusis/gochimp/mandrill/api"
)

// APIError represents an error returned by the Mandrill API
type APIError struct {
	status  string
	code    int
	name    string
	message string
	err     string
}

// InvalidKeyError is the error returned by Mandrill for invalid credentials
type InvalidKeyError APIError

// Error returns the error message
func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("endpoint returned an authentication error: %s", e.err)
}

// PaymentRequiredError is the error returned by Mandrill for payment related issues
type PaymentRequiredError APIError

// Error returns the error message
func (e *PaymentRequiredError) Error() string {
	return fmt.Sprintf("endpoint returned a payment related error: %s", e.err)
}

// SubAccountError is the error returned when a request is made with an invalid subaccount
type SubAccountError APIError

// Error returns the error message
func (e *SubAccountError) Error() string {
	return fmt.Sprintf("the provided subaccount was not valid: %s", e.err)
}

// TemplateError is the error returned when a template is invalid
type TemplateError APIError

// Error returns the error message
func (e *TemplateError) Error() string {
	return fmt.Sprintf("the provided template did not exist: %s", e.err)
}

// ServiceUnavailbleError is the error returned when a required subservice is unavailble for an api call (i.e. search)
type ServiceUnavailableError APIError

// Error returns the error message
func (e *ServiceUnavailableError) Error() string {
	return fmt.Sprintf("a required subsystem is down for maintenance: %s", e.err)
}

// ValidationError is the error returned when some aspect of the post data in the api call is invalid
type ValidationError APIError

// Error returns the error message
func (e *ValidationError) Error() string {
	return fmt.Sprintf("endpoint returned a validation error: %s", e.err)
}

// GeneralError is the Mandrill error for all other errors
type GeneralError APIError

// Error returns the error message
func (e *GeneralError) Error() string {
	return fmt.Sprintf("endpoint returned an unknown error: %s", e.err)
}

// UnknownError is an error type for Mandrill errors that we are unable to identify
type UnknownError APIError

// Error returns the error message
func (e *UnknownError) Error() string {
	return fmt.Sprintf("unknown error | message: %s | code: %d | name: %s | status: %s", e.message, e.code, e.name, e.status)
}

// JSONError is a custom error type for indicating JSON encode/decode issues
type JSONError struct {
	err string
}

// Error returns the error message
func (e *JSONError) Error() string {
	return fmt.Sprintf("error parsing json: %s", e.err)
}

func parseErrorJSON(b []byte) error {
	var e api.ErrorResponse
	// we're going to implement a strict decoder here
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&e); err != nil {
		// response isn't a mandrill error we can parse so just return nil
		// and let caller bubble up original error
		return nil
	}
	switch e.Name {
	case "Invalid_Key":
		return &InvalidKeyError{err: e.Message, status: e.Status, code: e.Code}
	case "Unknown_Template":
		return &TemplateError{err: e.Message, status: e.Status, code: e.Code}
	case "PaymentRequired":
		return &PaymentRequiredError{err: e.Message, status: e.Status, code: e.Code}
	case "Unknown_Subaccount":
		return &SubAccountError{err: e.Message, status: e.Status, code: e.Code}
	case "ValidationError":
		return &ValidationError{err: e.Message, status: e.Status, code: e.Code}
	case "GeneralError":
		return &GeneralError{err: e.Message, status: e.Status, code: e.Code}
	case "ServiceUnavailable":
		return &ServiceUnavailableError{err: e.Message, status: e.Status, code: e.Code}
	default:
		return &UnknownError{message: e.Message, status: e.Status, code: e.Code, name: e.Name}
	}
}
