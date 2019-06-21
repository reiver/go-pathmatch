package pathmatch

import (
	"sync"
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
	mutex sync.RWMutex
	bits              []string
	names             []string
	namesSet map[string]struct{}
	fieldTagName        string
}

func (pattern *Pattern) MatchNames() []string {
	if nil == pattern {
		return nil
	}

	pattern.mutex.RLock()
	defer pattern.mutex.RUnlock()

	return pattern.names
}
