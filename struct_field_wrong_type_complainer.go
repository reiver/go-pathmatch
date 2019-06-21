package pathmatch

import (
	"bytes"
	"fmt"
)

const structFieldWrongTypeMessagePrefix = "Bad Request: Wrong type for match "

type StructFieldWrongType interface {
	BadRequestComplainer
	MatchName() string
}

// internalStructFieldWrongType is the only underlying implementation that fits the
// StructFieldWrongType interface, in this library.
type internalStructFieldWrongType struct {
	matchName string
}

// newStructFieldWrongType creates a new internalStructFieldWrongType (struct) and
// returns it as a StructFieldWrongType (interface).
func newStructFieldWrongType(matchName string) StructFieldWrongType {
	err := internalStructFieldWrongType{
		matchName:matchName,
	}

	return &err
}

// Error method is necessary to satisfy the 'error' interface (and the StructFieldWrongType
// interface).
func (err *internalStructFieldWrongType) Error() string {
	var buffer bytes.Buffer

	buffer.WriteString(structFieldWrongTypeMessagePrefix)
	buffer.WriteString(fmt.Sprintf("%q", err.matchName))

	return buffer.String()
}

// BadRequestComplainer method is necessary to satisfy the 'BadRequestComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalStructFieldWrongType) BadRequestComplainer() {
        // Nothing here.
}

// DependencyName method is necessary to satisfy the 'StructFieldWrongType' interface.
func (err *internalStructFieldWrongType) MatchName() string {
	return err.matchName
}
