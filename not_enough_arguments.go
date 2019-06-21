package pathmatch

import (
	"fmt"
)

type NotEnoughArguments interface {
	BadRequest
	NotEnoughArguments()

	ExpectedAtLeast() int
	Actual() int
}

// internalNotEnoughArguments is the only underlying implementation that fits the
// NotEnoughArguments interface, in this library.
type internalNotEnoughArguments struct {
	expectedAtLeast int
	actual          int
}

// newNotEnoughArguments creates a new internalNotEnoughArguments (struct) and
// returns it as a NotEnoughArguments (interface).
func newNotEnoughArguments(expectedAtLeast int, actual int) NotEnoughArguments {
	err := internalNotEnoughArguments{
		expectedAtLeast:expectedAtLeast,
		actual:actual,
	}

	return &err
}

// Error method is necessary to satisfy the 'error' interface (and the
// NotEnoughArguments interface).
func (err *internalNotEnoughArguments) Error() string {
	plural := ""
	if 1 < err.expectedAtLeast {
		plural = "s"
	}
	return fmt.Sprintf("Bad Request: Not enough arguments. Expected at least %d argument%s, but actually got %d.", err.expectedAtLeast, plural, err.actual)
}

// BadRequest method is necessary to satisfy the 'BadRequest' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalNotEnoughArguments) BadRequest() {
	// Nothing here.
}

// NotEnoughArguments method is necessary to satisfy the 'NotEnoughArguments' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalNotEnoughArguments) NotEnoughArguments() {
	// Nothing here.
}

func (err *internalNotEnoughArguments) ExpectedAtLeast() int {
	return err.expectedAtLeast
}

func (err *internalNotEnoughArguments) Actual() int {
	return err.actual
}
