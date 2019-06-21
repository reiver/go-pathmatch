package pathmatch

import (
	"fmt"
)

type ScanError interface {
	InternalError
	ScanError()
	WrappedError() error
}

// internalScanError is the only underlying implementation that fits the
// ScanError interface, in this library.
type internalScanError struct {
	wrappedError error
	argumentIndex int
	argumentType  string
}

// newScanError creates a new internalScanError (struct) and
// returns it as a ScanError (interface).
func newScanError(wrappedError error, argumentIndex int, argumentType string) ScanError {
	err := internalScanError{
		wrappedError:wrappedError,
		argumentIndex:argumentIndex,
		argumentType:argumentType,
	}

	return &err
}

// Error method is necessary to satisfy the 'error' interface (and the
// ScanError interface).
func (err *internalScanError) Error() string {
	s := fmt.Sprintf("Internal Error: Received scan error for argument #%d (%s): %q", err.argumentIndex, err.argumentType, err.wrappedError.Error())
	return s
}

// InternalError method is necessary to satisfy the 'InternalError' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalScanError) InternalError() {
	// Nothing here.
}

// ScanError method is necessary to satisfy the 'ScanError' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalScanError) ScanError() {
	// Nothing here.
}

func (err *internalScanError) WrappedError() error {
	return err.wrappedError
}
