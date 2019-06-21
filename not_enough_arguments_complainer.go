package pathmatch


import (
	"fmt"
)


type NotEnoughArgumentsComplainer interface {
	BadRequest
	NotEnoughArgumentsComplainer()

	ExpectedAtLeast() int
	Actual() int
}


// internalNotEnoughArgumentsComplainer is the only underlying implementation that fits the
// NotEnoughArgumentsComplainer interface, in this library.
type internalNotEnoughArgumentsComplainer struct {
	expectedAtLeast int
	actual          int
}


// newNotEnoughArgumentsComplainer creates a new internalNotEnoughArgumentsComplainer (struct) and
// returns it as a NotEnoughArgumentsComplainer (interface).
func newNotEnoughArgumentsComplainer(expectedAtLeast int, actual int) NotEnoughArgumentsComplainer {
	err := internalNotEnoughArgumentsComplainer{
		expectedAtLeast:expectedAtLeast,
		actual:actual,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the
// NotEnoughArgumentsComplainer interface).
func (err *internalNotEnoughArgumentsComplainer) Error() string {
	plural := ""
	if 1 < err.expectedAtLeast {
		plural = "s"
	}
	return fmt.Sprintf("Bad Request: Not enough arguments. Expected at least %d argument%s, but actually got %d.", err.expectedAtLeast, plural, err.actual)
}


// BadRequest method is necessary to satisfy the 'BadRequest' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalNotEnoughArgumentsComplainer) BadRequest() {
	// Nothing here.
}


// NotEnoughArgumentsComplainer method is necessary to satisfy the 'NotEnoughArgumentsComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalNotEnoughArgumentsComplainer) NotEnoughArgumentsComplainer() {
	// Nothing here.
}


func (err *internalNotEnoughArgumentsComplainer) ExpectedAtLeast() int {
	return err.expectedAtLeast
}

func (err *internalNotEnoughArgumentsComplainer) Actual() int {
	return err.actual
}
