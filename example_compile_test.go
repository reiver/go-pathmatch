package pathmatch_test

import (
	"github.com/reiver/go-pathmatch"

	"fmt"
)

func ExampleCompile() {

	pattern, err := pathmatch.Compile("/v1/users/{user_id}/contacts/{contact_type}")
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	target := struct{
		UserID      string `match:"user_id"`
		ContactType string `match:"contact_type"`
	}{}


	var path = "/v1/users/123/contacts/e-mail"

	matched, err := pattern.MatchAndLoad(path, &target)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	if !matched {
		fmt.Printf("The path %q did not match.", path)
		return
	}

	fmt.Printf("user_id      = %q\n", target.UserID)
	fmt.Printf("contact_type = %q\n", target.ContactType)

	// Output:
	// user_id      = "123"
	// contact_type = "e-mail"
}
