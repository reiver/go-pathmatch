package pathmatch


import (
	"fmt"
)


type ScanErrorComplainer interface {
	InternalErrorComplainer
	ScanErrorComplainer()
	WrappedError() error
}


// internalScanErrorComplainer is the only underlying implementation that fits the
// ScanErrorComplainer interface, in this library.
type internalScanErrorComplainer struct {
	wrappedError error
	argumentIndex int
	argumentType  string
}


// newScanErrorComplainer creates a new internalScanErrorComplainer (struct) and
// returns it as a ScanErrorComplainer (interface).
func newScanErrorComplainer(wrappedError error, argumentIndex int, argumentType string) ScanErrorComplainer {
	err := internalScanErrorComplainer{
		wrappedError:wrappedError,
		argumentIndex:argumentIndex,
		argumentType:argumentType,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the
// ScanErrorComplainer interface).
func (err *internalScanErrorComplainer) Error() string {
	s := fmt.Sprintf("Internal Error: Received scan error for argument #%d (%s): %q", err.argumentIndex, err.argumentType, err.wrappedError.Error())
	return s
}


// InternalErrorComplainer method is necessary to satisfy the 'InternalErrorComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalScanErrorComplainer) InternalErrorComplainer() {
	// Nothing here.
}


// ScanErrorComplainer method is necessary to satisfy the 'ScanErrorComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalScanErrorComplainer) ScanErrorComplainer() {
	// Nothing here.
}

func (err *internalScanErrorComplainer) WrappedError() error {
	return err.wrappedError
}
