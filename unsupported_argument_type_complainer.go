package pathmatch

import (
	"fmt"
)

type UnsupportedArgumentType interface {
	BadRequestComplainer
	UnsupportedArgumentType()
}

// internalUnsupportedIndexedArgumentTypeComplainer is the only underlying implementation that fits the
// UnsupportedArgumentType interface, in this library.
type internalUnsupportedIndexedArgumentTypeComplainer struct {
	argumentIndex int
	argumentType  string
}

// newUnsupportedIndexedArgumentTypeComplainer creates a new internalUnsupportedIndexedArgumentTypeComplainer (struct) and
// returns it as a UnsupportedArgumentType (interface).
func newUnsupportedIndexedArgumentTypeComplainer(argumentIndex int, argumentType string) UnsupportedArgumentType {
	err := internalUnsupportedIndexedArgumentTypeComplainer{
		argumentIndex:argumentIndex,
		argumentType:argumentType,
	}

	return &err
}

// Error method is necessary to satisfy the 'error' interface (and the
// UnsupportedArgumentType interface).
func (err *internalUnsupportedIndexedArgumentTypeComplainer) Error() string {
	s := fmt.Sprintf("Bad Request: Type of argument #%d (%s) is unsupported.", err.argumentIndex, err.argumentType)
	if "string" == err.argumentType {
		s = fmt.Sprintf("%s However, type \"*string\" (pointer to string) is supported; did you mean to use a \"*string\" instead?", s)
	}
	return s
}

// BadRequestComplainer method is necessary to satisfy the 'BadRequestComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalUnsupportedIndexedArgumentTypeComplainer) BadRequestComplainer() {
	// Nothing here.
}

// UnsupportedArgumentType method is necessary to satisfy the 'UnsupportedArgumentType' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalUnsupportedIndexedArgumentTypeComplainer) UnsupportedArgumentType() {
	// Nothing here.
}

// internalUnsupportedArgumentType is the only underlying implementation that fits the
// UnsupportedArgumentType interface, in this library.
type internalUnsupportedArgumentType struct {
	msg string
}

// newUnsupportedArgumentType creates a new internalUnsupportedArgumentType (struct) and
// returns it as a UnsupportedArgumentType (interface).
func newUnsupportedArgumentType(format string, a ...interface{}) UnsupportedArgumentType {
	msg := fmt.Sprintf(format, a...)

	err := internalUnsupportedArgumentType{
		msg:msg,
	}

	return &err
}

// Error method is necessary to satisfy the 'error' interface (and the
// UnsupportedArgumentType interface).
func (err *internalUnsupportedArgumentType) Error() string {

	return fmt.Sprintf("Bad Request: Unsupported Argument Type: %s", err.msg)
}

// BadRequestComplainer method is necessary to satisfy the 'BadRequestComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalUnsupportedArgumentType) BadRequestComplainer() {
	// Nothing here.
}

// UnsupportedArgumentType method is necessary to satisfy the 'UnsupportedArgumentType' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalUnsupportedArgumentType) UnsupportedArgumentType() {
	// Nothing here.
}
