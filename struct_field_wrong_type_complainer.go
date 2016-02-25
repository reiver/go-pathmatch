package pathmatch


import (
	"bytes"
	"fmt"
)


const structFieldWrongTypeMessagePrefix = "Bad Request: Wrong type for match "


type StructFieldWrongTypeComplainer interface {
	BadRequestComplainer
	MatchName() string
}

// internalStructFieldWrongTypeComplainer is the only underlying implementation that fits the
// StructFieldWrongTypeComplainer interface, in this library.
type internalStructFieldWrongTypeComplainer struct {
	matchName string
}

// newStructFieldWrongTypeComplainer creates a new internalStructFieldWrongTypeComplainer (struct) and
// returns it as a StructFieldWrongTypeComplainer (interface).
func newStructFieldWrongTypeComplainer(matchName string) StructFieldWrongTypeComplainer {
	err := internalStructFieldWrongTypeComplainer{
		matchName:matchName,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the StructFieldWrongTypeComplainer
// interface).
func (err *internalStructFieldWrongTypeComplainer) Error() string {
	var buffer bytes.Buffer

	buffer.WriteString(structFieldWrongTypeMessagePrefix)
	buffer.WriteString(fmt.Sprintf("%q", err.matchName))

	return buffer.String()
}


// BadRequestComplainer method is necessary to satisfy the 'BadRequestComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalStructFieldWrongTypeComplainer) BadRequestComplainer() {
        // Nothing here.
}


// DependencyName method is necessary to satisfy the 'StructFieldWrongTypeComplainer' interface.
func (err *internalStructFieldWrongTypeComplainer) MatchName() string {
	return err.matchName
}
