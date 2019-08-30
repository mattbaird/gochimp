package mandrill

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lusis/gochimp/mandrill/api"
)

type APIError struct {
	status  string
	code    int
	name    string
	message string
	err     string
}

type InvalidKeyError APIError

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("endpoint returned an authentication error: %s", e.err)
}

type PaymentRequiredError APIError

func (e *PaymentRequiredError) Error() string {
	return fmt.Sprintf("endpoint returned a payment related error: %s", e.err)
}

type SubAccountError APIError

func (e *SubAccountError) Error() string {
	return fmt.Sprintf("the provided subaccount was not valid: %s", e.err)
}

type TemplateError APIError

func (e *TemplateError) Error() string {
	return fmt.Sprintf("the provided template did not exist: %s", e.err)
}

type ServiceUnavailableError APIError

func (e *ServiceUnavailableError) Error() string {
	return fmt.Sprintf("a required subsystem is down for maintenance: %s", e.err)
}

type ValidationError APIError

func (e *ValidationError) Error() string {
	return fmt.Sprintf("endpoint returned a validation error: %s", e.err)
}

type GeneralError APIError

func (e *GeneralError) Error() string {
	return fmt.Sprintf("endpoint returned an unknown error: %s", e.err)
}

type UnknownError APIError

func (e *UnknownError) Error() string {
	return fmt.Sprintf("unknown error | message: %s | code: %d | name: %s | status: %s", e.message, e.code, e.name, e.status)
}

type JSONError struct {
	err string
}

func (e *JSONError) Error() string {
	return fmt.Sprintf("error parsing json: %s", e.err)
}

func parseErrorJSON(b []byte) error {
	var e api.ErrorResponse
	// we're going to implement a strict decoder here
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&e); err != nil {
		// response isn't a mandrill error we can parse so just bubble it up as is
		return err
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
