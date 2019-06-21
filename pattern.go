package pathmatch

import (
	"bytes"
)

// Pattern represents a compiled pattern. It is what is returned
// from calling either the Compile to MustCompile funcs.
//
// Pattern provides the Match, FindAndLoad, and MatchNames methods.
//
// Example Usage:
//
//	pattern, err := pathmath.Compile("/users/{user_id}")
//	if nil != err {
//		fmt.Printf("ERROR Compiling: %v\n", err)
//		return
//	}
//	
//	var userId string
//	
//	didMatch, err := pattern.Find("/users/123", userId)
//	if nil != err {
//		fmt.Printf("ERROR Matching: %v\n", err)
//		return
//	}
//	
//	if didMatch {
//		fmt.Printf("user_id = %q\n", userId)
//	} else {
//		fmt.Println("Did not match.")
//	}
type Pattern struct {
	bits              []string
	names             []string
	namesSet map[string]struct{}
	fieldTagName        string
}

func newPattern(fieldTagName string) *Pattern {
	bits     := []string{}
	names    := []string{}
	namesSet := map[string]struct{}{}

	pattern := Pattern{
		bits:bits,
		names:names,
		namesSet:namesSet,
		fieldTagName:fieldTagName,
	}

	return &pattern

}

func (pattern *Pattern) MatchNames() []string {

	return pattern.names
}

func (pattern *Pattern) Glob() string {
//@TODO: This shouldn't be executed every time!

	var buffer bytes.Buffer

	for _, bit := range pattern.bits {
		if wildcardBit == bit {
			buffer.WriteRune('*')
		} else {
			buffer.WriteString(bit)
		}
	}

	return buffer.String()
}
