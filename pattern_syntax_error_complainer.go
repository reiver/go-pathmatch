package pathmatch


import (
	"fmt"
)


// PatternSyntaxErrorComplainer is used to represent a specific kind of BadRequestComplainer error.
// Specifically, it represents a syntax error in the uncompiled pattern passed to the
// pathmatch.Compile() func.
//
// Example usage is as follows:
//
//	pattern, err := pathmatch.Compile("/something/{there_is_a_syntax_error_in_this_pattern")
//	if nil != err {
//		switch err.(type) {
//	
//		case pathmatch.PatternSyntaxErrorComplainer: // ← Here we are detecting if the error returned was due to a syntax error, in the uncompiled pattern. Also note that it comes BEFORE the 'pathmatch.BadRequestComplainer' case; THAT IS IMPORTANT!
//	
//			fmt.Printf("The uncompiled pattern passed to pathmatch.Compile() had a syntax error in it. The error message describing the syntax error is....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.BadRequestComplainer:
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalErrorComplainer:
//	
//			fmt.Printf("It's not your fault; it's our fault. Something bad happened internally when pathmatch.Compile() was running. The error message was....\n%s\n", err.Error())
//			return
//	
//		default:
//	
//			fmt.Printf("Some kind of unexpected error happend: %v", err)
//			return
//		}
//	}
type PatternSyntaxErrorComplainer interface {
	BadRequestComplainer
	PatternSyntaxErrorComplainer()
}


// internalPatternSyntaxErrorComplainer is the only underlying implementation that fits the
// PatternSyntaxErrorComplainer interface, in this library.
type internalPatternSyntaxErrorComplainer struct {
	msg string
}


// newPatternSyntaxErrorComplainer creates a new internalPatternSyntaxErrorComplainer (struct) and
// returns it as a PatternSyntaxErrorComplainer (interface).
func newPatternSyntaxErrorComplainer(format string, a ...interface{}) PatternSyntaxErrorComplainer {
	msg := fmt.Sprintf(format, a...)

	err := internalPatternSyntaxErrorComplainer{
		msg:msg,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the
// PatternSyntaxErrorComplainer interface).
func (err *internalPatternSyntaxErrorComplainer) Error() string {
	s := fmt.Sprintf("Bad Request: Syntax Error: %s", err.msg)
	return s
}


// BadRequestComplainer method is necessary to satisfy the 'InternalErrorComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalPatternSyntaxErrorComplainer) BadRequestComplainer() {
	// Nothing here.
}


// PatternSyntaxErrorComplainer method is necessary to satisfy the 'PatternSyntaxErrorComplainer' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalPatternSyntaxErrorComplainer) PatternSyntaxErrorComplainer() {
	// Nothing here.
}
