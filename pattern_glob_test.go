package pathmatch


import (
	"testing"
)


func TestGlob(t *testing.T) {

	tests := []struct{
		Pattern  string
		Expected string
	}{
		{
			Pattern:  "/",
			Expected: "/",
		},
		{
			Pattern:  "/apple",
			Expected: "/apple",
		},
		{
			Pattern:  "/apple/",
			Expected: "/apple/",
		},
		{
			Pattern:  "/apple/banana",
			Expected: "/apple/banana",
		},
		{
			Pattern:  "/apple/banana/",
			Expected: "/apple/banana/",
		},
		{
			Pattern:  "/apple/banana/cherry",
			Expected: "/apple/banana/cherry",
		},
		{
			Pattern:  "/apple/banana/cherry/",
			Expected: "/apple/banana/cherry/",
		},



		{
			Pattern:  "",
			Expected: "",
		},
		{
			Pattern:  "apple",
			Expected: "apple",
		},
		{
			Pattern:  "apple/",
			Expected: "apple/",
		},
		{
			Pattern:  "apple/banana",
			Expected: "apple/banana",
		},
		{
			Pattern:  "apple/banana/",
			Expected: "apple/banana/",
		},
		{
			Pattern:  "apple/banana/cherry",
			Expected: "apple/banana/cherry",
		},
		{
			Pattern:  "apple/banana/cherry/",
			Expected: "apple/banana/cherry/",
		},



		{
			Pattern:  "/users/{user_id}",
			Expected: "/users/*",
		},
		{
			Pattern:  "/users/{user_id}/",
			Expected: "/users/*/",
		},
		{
			Pattern:  "/users/{user_id}/cards",
			Expected: "/users/*/cards",
		},
		{
			Pattern:  "/users/{user_id}/cards/{card_id}",
			Expected: "/users/*/cards/*",
		},
		{
			Pattern:  "/users/{user_id}/cards/{card_id}/",
			Expected: "/users/*/cards/*/",
		},



		{
			Pattern:  "users/{user_id}",
			Expected: "users/*",
		},
		{
			Pattern:  "users/{user_id}/",
			Expected: "users/*/",
		},
		{
			Pattern:  "users/{user_id}/cards",
			Expected: "users/*/cards",
		},
		{
			Pattern:  "users/{user_id}/cards/{card_id}",
			Expected: "users/*/cards/*",
		},
		{
			Pattern:  "users/{user_id}/cards/{card_id}/",
			Expected: "users/*/cards/*/",
		},



		{
			Pattern:  "/{this}",
			Expected: "/*",
		},
		{
			Pattern:  "/{this}/",
			Expected: "/*/",
		},
		{
			Pattern:  "/{this}/{that}",
			Expected: "/*/*",
		},
		{
			Pattern:  "/{this}/{that}/",
			Expected: "/*/*/",
		},
		{
			Pattern:  "/{this}/{that}/{these}",
			Expected: "/*/*/*",
		},
		{
			Pattern:  "/{this}/{that}/{these}/",
			Expected: "/*/*/*/",
		},
		{
			Pattern:  "/{this}/{that}/{these}/{those}",
			Expected: "/*/*/*/*",
		},
		{
			Pattern:  "/{this}/{that}/{these}/{those}/",
			Expected: "/*/*/*/*/",
		},



		{
			Pattern:  "{this}",
			Expected: "*",
		},
		{
			Pattern:  "{this}/",
			Expected: "*/",
		},
		{
			Pattern:  "{this}/{that}",
			Expected: "*/*",
		},
		{
			Pattern:  "{this}/{that}/",
			Expected: "*/*/",
		},
		{
			Pattern:  "{this}/{that}/{these}",
			Expected: "*/*/*",
		},
		{
			Pattern:  "{this}/{that}/{these}/",
			Expected: "*/*/*/",
		},
		{
			Pattern:  "{this}/{that}/{these}/{those}",
			Expected: "*/*/*/*",
		},
		{
			Pattern:  "{this}/{that}/{these}/{those}/",
			Expected: "*/*/*/*/",
		},
	}


	for testNumber, test := range tests {

		pattern, err := Compile(test.Pattern)
		if nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: %v", testNumber, err)
			continue
		}

		glob := pattern.Glob()
		if expected, actual := test.Expected, glob; expected != actual {
			t.Errorf("For test #%d, expected glob to be %q, but was actually %q.", testNumber, expected, actual)
			continue
		}
	}
}
