package events

import (
	"fmt"
)

// ErrInvalidSignature is the error for when the signature is invalid
type ErrInvalidSignature struct {
	computed string
	provided string
}

// Error returns the error string
func (e ErrInvalidSignature) Error() string {
	return fmt.Sprintf("invalid signature. provided: %s computed: %s", e.provided, e.computed)
}

// ErrNoSignature is the error type for missing mandrill signature header
type ErrNoSignature struct{}

// Error returns the error string
func (e ErrNoSignature) Error() string {
	return "missing the X-MANDRILL-SIGNATURE header"
}

// ErrNotPost is the error type for invalid http method in webhook
type ErrNotPost struct {
	method string
}

// Error returns the error string
func (e ErrNotPost) Error() string {
	return fmt.Sprintf("webhook is not a POST request. Got %s", e.method)
}

// ErrParse is the error type for invalid postform data in webhook
type ErrParse struct {
	parseErr string
}

// Error returns the error string
func (e ErrParse) Error() string {
	return fmt.Sprintf("unable to parse webhook: %s", e.parseErr)
}

// ErrMissingMandrillEvents is the error type for missing mandrill_events form data
type ErrMissingMandrillEvents struct {
	postForm string
}

// Error returns the error string
func (e ErrMissingMandrillEvents) Error() string {
	return fmt.Sprintf("mandrill_events missing in webhook: Got %s", e.postForm)
}

// ErrCryptoErr is the error type for issues related to crypto in signature verification
type ErrCryptoErr struct {
	err string
}

// Error returns the error string
func (e ErrCryptoErr) Error() string {
	return fmt.Sprintf("error generating hash: %s", e.err)
}

// UnmarshallError is an error when attempting to decode json
type UnmarshallError struct {
	data string
	msg  string
}

// Error returns the error string
func (e UnmarshallError) Error() string {
	return fmt.Sprintf("error decoding json: %s (data: %s)", e.msg, e.data)
}

// InvalidEventType represents an unhandled event type
type InvalidEventType struct {
	eventType string
}

// Error returns the error message
func (e InvalidEventType) Error() string {
	return fmt.Sprintf("unknown event type %s", e.eventType)
}
