package pathmatch

import (
	"bytes"
	"fmt"
)

const (
	internalErrorMessagePrefix = "Internal Error: "
)

type InternalError interface {
	error
	InternalError()
}

// internalInternalError is the only underlying implementation that fits the
// InternalError interface, in this library.
type internalInternalError struct {
	msg string
}

// newInternalError creates a new internalInternalError (struct) and
// returns it as a InternalError (interface).
func newInternalError(format string, a ...interface{}) InternalError {
	msg := fmt.Sprintf(format, a...)

	err := internalInternalError{
		msg:msg,
	}

	return &err
}

// Error method is necessary to satisfy the 'error' interface (and the InternalError
// interface).
func (err *internalInternalError) Error() string {
	var buffer bytes.Buffer

	buffer.WriteString(internalErrorMessagePrefix)
	buffer.WriteString(err.msg)

	return buffer.String()
}

// InternalError method is necessary to satisfy the 'InternalError' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalInternalError) InternalError() {
	// Nothing here.
}
