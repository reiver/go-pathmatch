package pathmatch


import (
	"strings"
)


const (
	defaultFieldTagName = "match"
	wildcardBit = "{}"
)


var (
	errMissingEndingRightBraceToMatchBeginningLeftBrace = newPatternSyntaxError(`Missing ending "}" (to match beginning "{").`)
	errSlashInsideOfBraces     = newPatternSyntaxError(`"/" inside of "{...}".`)
	errLeftBraceInsideOfBraces = newPatternSyntaxError(`"{" inside of "{...}".`)
)


// Compile takes an uncompiled pattern, in the form of a Go string (ex: "/users/{userId}/vehicles/{vehicleId}"),
// and returns a compiled pattern.
//
// The compiled pattern can then be used to test if a path matches the pattern it contains.
//
// If the uncompiled pattern has a syntax error, Compile returns an error.
//
// Example Usage:
//
//	pattern, err := pathmath.Compile("/users/{user_id}")
//	if nil != err {
//		fmt.Printf("ERROR Compiling: %v\n", err)
//		return
//	}
func Compile(uncompiledPattern string) (*Pattern, error) {

	var pattern Pattern

	err := CompileTo(&pattern, uncompiledPattern)
	if nil != err {
		return nil, err
	}

	return &pattern, nil
}

// CompileTo takes an uncompiled pattern, in the form of a Go string (ex: "/users/{userId}/vehicles/{vehicleId}"),
// and compiles the pattern to the ‘target’.
//
// The compiled pattern can then be used to test if a path matches the pattern it contains.
//
// If the uncompiled pattern has a syntax error, Compile returns an error.
//
// Example Usage:
//
//	var pattern patchmatch.Pattern
//
//	err := pathmath.CompileTo(&pattern, "/users/{user_id}")
//	if nil != err {
//		fmt.Printf("ERROR Compiling: %v\n", err)
//		return
//	}
func CompileTo(target *Pattern, uncompiledPattern string) error {
	if nil == target {
		return errNilTarget
	}

	newPattern(target, defaultFieldTagName)

	s := uncompiledPattern
	for {
		index := strings.IndexRune(s, '{')
		if -1 == index {
			target.bits = append(target.bits, s)
			break
		}
		bit := s[:index]
		if "" != bit { // This is to deal with the case where a {???} is right at the beginning of the uncompiledPattern.
			target.bits = append(target.bits, bit)
		}
		s = s[1+index:]
		if "" == s {
			break
		}

		index = strings.IndexRune(s, '}')
		if -1 == index {
			return errMissingEndingRightBraceToMatchBeginningLeftBrace
		}

		// There should not be a slash ("/") before the ending brace ("}").
		// If there is, it is a syntax error.
		slashIndex := strings.IndexRune(s, '/')
		if -1 != slashIndex && slashIndex <= index {
			return errSlashInsideOfBraces
		}

		// There should not be another beginning brace ("{") before the ending brace ("}").
		// If there is, it is a syntax error.
		anotherLeftBraceIndex := strings.IndexRune(s, '{')
		if -1 != anotherLeftBraceIndex && anotherLeftBraceIndex <= index {
			return errLeftBraceInsideOfBraces
		}


		bit = s[:index]


		// Match names should be unique, within a pattern.
		if _, ok := target.namesSet[bit]; ok {
			return newPatternSyntaxError("Duplicate match name: %q.", bit)
		}


		target.names = append(target.names, bit)
		target.namesSet[bit] = struct{}{}
		target.bits  = append(target.bits, wildcardBit)
		s = s[1+index:]
		if "" == s {
			break
		}
	}


	return nil
}


// MustCompile is like Compile except that it never returns an error; but
// instead panic()s if there was an error.
//
// Example Usage:
//
//	pattern := pathmath.MustCompile("/users/{user_id}")
//
// Note that if one recover()s from the panic(), one can use a Go type-switch
// to figure out what kind of error it is.
func MustCompile(uncompiledPattern string) *Pattern {
	var pattern Pattern

	if err := CompileTo(&pattern, uncompiledPattern); nil != err {
		panic(err)
	} else {
		return &pattern
	}
}
