package pathmatch


import (
	"fmt"
)


// PatternSyntaxError is used to represent a specific kind of BadRequest error.
// Specifically, it represents a syntax error in the uncompiled pattern passed to the
// pathmatch.Compile() func.
//
// Example usage is as follows:
//
//	pattern, err := pathmatch.Compile("/something/{there_is_a_syntax_error_in_this_pattern")
//	if nil != err {
//		switch err.(type) {
//	
//		case pathmatch.PatternSyntaxError: // ← Here we are detecting if the error returned was due to a syntax error, in the uncompiled pattern. Also note that it comes BEFORE the 'pathmatch.BadRequest' case; THAT IS IMPORTANT!
//	
//			fmt.Printf("The uncompiled pattern passed to pathmatch.Compile() had a syntax error in it. The error message describing the syntax error is....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.BadRequest:
//	
//			fmt.Printf("Something you did when you called pathmatch.Compile() caused an error. The error message was....\n%s\n", err.Error())
//			return
//	
//		case pathmatch.InternalError:
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
type PatternSyntaxError interface {
	BadRequest
	PatternSyntaxError()
}


// internalPatternSyntaxError is the only underlying implementation that fits the
// PatternSyntaxError interface, in this library.
type internalPatternSyntaxError struct {
	msg string
}


// newPatternSyntaxError creates a new internalPatternSyntaxError (struct) and
// returns it as a PatternSyntaxError (interface).
func newPatternSyntaxError(format string, a ...interface{}) PatternSyntaxError {
	msg := fmt.Sprintf(format, a...)

	err := internalPatternSyntaxError{
		msg:msg,
	}

	return &err
}


// Error method is necessary to satisfy the 'error' interface (and the
// PatternSyntaxError interface).
func (err *internalPatternSyntaxError) Error() string {
	s := fmt.Sprintf("Bad Request: Syntax Error: %s", err.msg)
	return s
}


// BadRequest method is necessary to satisfy the 'InternalError' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalPatternSyntaxError) BadRequest() {
	// Nothing here.
}


// PatternSyntaxError method is necessary to satisfy the 'PatternSyntaxError' interface.
// It exists to make this error type detectable in a Go type-switch.
func (err *internalPatternSyntaxError) PatternSyntaxError() {
	// Nothing here.
}
