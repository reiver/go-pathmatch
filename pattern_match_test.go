package pathmatch


import (
	"testing"
)


func TestMatch(t *testing.T) {

	tests := []struct{
		Pattern       Pattern
		Args         []interface{}
		Path           string
		ExpectedArgs []string
	}{
		{
			Pattern: MustCompile("/{this}/{that}/{these}/{those}"),
			Args: []interface{}{new(string), new(string), new(string), new(string), },
			Path:                 "/apple/banana/cherry/grape",
			ExpectedArgs: []string{"apple","banana","cherry","grape"},
		},



		{
			Pattern: MustCompile("/user/{sessionKey}"),
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/"),
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},

		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle"),
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/"),
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},

		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/DEFAULT"),
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/DEFAULT/"),
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},

		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/{vehicleIdcode}"),
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "DEFAULT"},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/{vehicleIdcode}/"),
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "DEFAULT"},
		},

		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/{vehicleIdcode}"),
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/N9Z_tiv7",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "N9Z_tiv7"},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/{vehicleIdcode}/"),
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/N9Z_tiv7/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "N9Z_tiv7"},
		},
	}


	for testNumber, test := range tests {

		for argNumber, arg := range test.Args {
			argStringPtr, ok := arg.(*string)
			if !ok {
				t.Errorf("For test #%d, expected test.Args[%d] to be of type *string, but actually was %T.", testNumber, argNumber, arg)
				continue
			}

			if expected, actual := "", *argStringPtr; expected != actual {
				t.Errorf("For test #%d, expected *test.Args[%d] to (initially) be %q, but actually was %q.", testNumber, argNumber, expected, actual)
				continue
			}
		}

		if didMatch, err := test.Pattern.Match(test.Path, test.Args...); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: %v", testNumber, err)
			continue
		} else if !didMatch {
			t.Errorf("For test #%d, expected path to match pattern, but it didn't.", testNumber)
			continue
		}

		for argNumber, arg := range test.Args {
			argStringPtr, ok := arg.(*string)
			if !ok {
				t.Errorf("For test #%d, expected test.Args[%d] to be of type *string, but actually was %T.", testNumber, argNumber, arg)
				continue
			}

			if expected, actual := test.ExpectedArgs[argNumber], *argStringPtr; expected != actual {
				t.Errorf("For test #%d, expected *test.Args[%d] to be %q, but actually was %q.", testNumber, argNumber, expected, actual)
				continue
			}
		}
	}
}
