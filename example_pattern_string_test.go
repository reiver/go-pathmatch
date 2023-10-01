package pathmatch_test

import (
	"sourcecode.social/reiver/go-pathmatch"

	"fmt"
	"os"
)

func ExamplePattern_String() {

	var template = "/v1/users/{user_id}"

	var pattern pathmatch.Pattern

	err := pathmatch.CompileTo(&pattern, template)
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return
	}

	fmt.Printf("pattern: %s", pattern)

	// Output:
	// pattern: /v1/users/{user_id}
}
