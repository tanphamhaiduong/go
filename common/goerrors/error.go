package goerrors

import (
	"fmt"
)

var (
	// ErrMessageBadRquestError ...
	ErrMessageBadRquestError = "[E4000] Bad Request"
	// ErrMessageNotAuthorized ...
	ErrMessageNotAuthorized = "[E4010] Not Authorized"
	// ErrMessageForBidden ...
	ErrMessageForBidden = "[E4030] Forbidden"
	// ErrMessageNotFound ...
	ErrMessageNotFound = "[E4040] Not Found"
	// ErrMessageRequestTimeout ...
	ErrMessageRequestTimeout = "[E4080] Request Timeout"
	// ErrMessageInternalServerError ...
	ErrMessageInternalServerError = "[E5000] Internal Server Error"
	// ErrMessageGatewayTimeout ...
	ErrMessageGatewayTimeout = "[E5040] Gateway Timeout"
)

// ErrBadRquestError ...
func ErrBadRquestError(traceID interface{}) error {
	return fmt.Errorf("%s %s", ErrMessageBadRquestError, traceID)
}

// ErrNotAuthorized ...
func ErrNotAuthorized(traceID interface{}) error {
	return fmt.Errorf("%s %s", ErrMessageNotAuthorized, traceID)
}

// ErrForBidden ...
func ErrForBidden(traceID interface{}) error {
	return fmt.Errorf("%s %s", ErrMessageForBidden, traceID)
}

// ErrNotFound ...
func ErrNotFound(traceID interface{}) error {
	return fmt.Errorf("%s %s", ErrMessageNotFound, traceID)
}

// ErrRequestTimeout ...
func ErrRequestTimeout(traceID interface{}) error {
	return fmt.Errorf("%s %s", ErrMessageRequestTimeout, traceID)
}

// ErrInternalServerError ...
func ErrInternalServerError(traceID interface{}) error {
	return fmt.Errorf("%s %s", ErrMessageInternalServerError, traceID)
}

// ErrGatewayTimeout ...
func ErrGatewayTimeout(traceID interface{}) error {
	return fmt.Errorf("%s %s", ErrMessageInternalServerError, traceID)
}
