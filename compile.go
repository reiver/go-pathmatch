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

	newPattern(&pattern, defaultFieldTagName)

	s := uncompiledPattern
	for {
		index := strings.IndexRune(s, '{')
		if -1 == index {
			pattern.bits = append(pattern.bits, s)
			break
		}
		bit := s[:index]
		if "" != bit { // This is to deal with the case where a {???} is right at the beginning of the uncompiledPattern.
			pattern.bits = append(pattern.bits, bit)
		}
		s = s[1+index:]
		if "" == s {
			break
		}

		index = strings.IndexRune(s, '}')
		if -1 == index {
			return nil, errMissingEndingRightBraceToMatchBeginningLeftBrace
		}

		// There should not be a slash ("/") before the ending brace ("}").
		// If there is, it is a syntax error.
		slashIndex := strings.IndexRune(s, '/')
		if -1 != slashIndex && slashIndex <= index {
			return nil, errSlashInsideOfBraces
		}

		// There should not be another beginning brace ("{") before the ending brace ("}").
		// If there is, it is a syntax error.
		anotherLeftBraceIndex := strings.IndexRune(s, '{')
		if -1 != anotherLeftBraceIndex && anotherLeftBraceIndex <= index {
			return nil, errLeftBraceInsideOfBraces
		}


		bit = s[:index]


		// Match names should be unique, within a pattern.
		if _, ok := pattern.namesSet[bit]; ok {
			return nil, newPatternSyntaxError("Duplicate match name: %q.", bit)
		}


		pattern.names = append(pattern.names, bit)
		pattern.namesSet[bit] = struct{}{}
		pattern.bits  = append(pattern.bits, wildcardBit)
		s = s[1+index:]
		if "" == s {
			break
		}
	}


	return &pattern, nil
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
	if pattern, err := Compile(uncompiledPattern); nil != err {
		panic(err)
	} else {
		return pattern
	}
}
