package pathmatch


import (
	"fmt"
)


type UnsupportedArgumentTypeComplainer interface {
	BadRequestComplainer
	UnsupportedArgumentTypeComplainer()
}






// internalUnsupportedIndexedArgumentTypeComplainer is the only underlying implementation that fits the
// UnsupportedArgumentTypeComplainer interface, in this library.
type internalUnsupportedIndexedArgumentTypeComplainer struct {
	argumentIndex int
	argumentType  string
}


// newUnsupportedIndexedArgumentTypeComplainer creates a new internalUnsupportedIndexedArgumentTypeComplainer (struct) and
// returns it as a UnsupportedArgumentTypeComplainer (interface).
func newUnsupportedIndexedArgumentTypeComplainer(argumentIndex int, argumentType string) UnsupportedArgumentTypeComplainer {
	err := internalUnsupportedIndexedArgumentTypeComplainer{
		argumentIndex:argumentIndex,
		argumentType:argumentType,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the
// UnsupportedArgumentTypeComplainer interface).
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


// UnsupportedArgumentTypeComplainer method is necessary to satisfy the 'UnsupportedArgumentTypeComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalUnsupportedIndexedArgumentTypeComplainer) UnsupportedArgumentTypeComplainer() {
	// Nothing here.
}






// internalUnsupportedArgumentTypeComplainer is the only underlying implementation that fits the
// UnsupportedArgumentTypeComplainer interface, in this library.
type internalUnsupportedArgumentTypeComplainer struct {
	msg string
}


// newUnsupportedArgumentTypeComplainer creates a new internalUnsupportedArgumentTypeComplainer (struct) and
// returns it as a UnsupportedArgumentTypeComplainer (interface).
func newUnsupportedArgumentTypeComplainer(format string, a ...interface{}) UnsupportedArgumentTypeComplainer {
	msg := fmt.Sprintf(format, a...)

	err := internalUnsupportedArgumentTypeComplainer{
		msg:msg,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the
// UnsupportedArgumentTypeComplainer interface).
func (err *internalUnsupportedArgumentTypeComplainer) Error() string {

	return fmt.Sprintf("Bad Request: Unsupported Argument Type: %s", err.msg)
}


// BadRequestComplainer method is necessary to satisfy the 'BadRequestComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalUnsupportedArgumentTypeComplainer) BadRequestComplainer() {
	// Nothing here.
}


// UnsupportedArgumentTypeComplainer method is necessary to satisfy the 'UnsupportedArgumentTypeComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalUnsupportedArgumentTypeComplainer) UnsupportedArgumentTypeComplainer() {
	// Nothing here.
}
