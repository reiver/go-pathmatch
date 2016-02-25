package pathmatch


import (
	"bytes"
	"fmt"
)


const (
	internalErrorMessagePrefix = "Internal Error: "
)


type InternalErrorComplainer interface {
	error
	InternalErrorComplainer()
}


// internalInternalErrorComplainer is the only underlying implementation that fits the
// InternalErrorComplainer interface, in this library.
type internalInternalErrorComplainer struct {
	msg string
}


// newInternalErrorComplainer creates a new internalInternalErrorComplainer (struct) and
// returns it as a InternalErrorComplainer (interface).
func newInternalErrorComplainer(format string, a ...interface{}) InternalErrorComplainer {
	msg := fmt.Sprintf(format, a...)

	err := internalInternalErrorComplainer{
		msg:msg,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the InternalErrorComplainer
// interface).
func (err *internalInternalErrorComplainer) Error() string {
	var buffer bytes.Buffer

	buffer.WriteString(internalErrorMessagePrefix)
	buffer.WriteString(err.msg)

	return buffer.String()
}


// InternalErrorComplainer method is necessary to satisfy the 'InternalErrorComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalInternalErrorComplainer) InternalErrorComplainer() {
	// Nothing here.
}
