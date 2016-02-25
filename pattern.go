package pathmatch


// Pattern represents a compiled pattern. It is what is returned
// from calling either the Compile to MustCompile funcs.
//
// Pattern provides the Match, MatchAndLoad, and MatchNames methods.
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
//	didMatch, err := pattern.Match("/users/123", userId)
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
type Pattern interface {
	Match(string, ...interface{}) (bool, error)
	MatchAndLoad(string, interface{}) (bool, error)
	MatchNames() []string
}


type internalPattern struct {
	bits              []string
	names             []string
	namesSet map[string]struct{}
	fieldTagName        string
}


func (pattern *internalPattern) MatchNames() []string {

	return pattern.names
}


func newPattern(fieldTagName string) *internalPattern {
	bits     := []string{}
	names    := []string{}
	namesSet := map[string]struct{}{}

	pattern := internalPattern{
		bits:bits,
		names:names,
		namesSet:namesSet,
		fieldTagName:fieldTagName,
	}

	return &pattern

}
