package pathmatch_test


import (
	"github.com/reiver/go-pathmatch"

	"testing"
)


func TestFind(t *testing.T) {

	tests := []struct{
		Pattern        string
		Args         []interface{}
		Path           string
		ExpectedArgs []string
	}{
		{
			Pattern: "/{this}/{that}/{these}/{those}",
			Args: []interface{}{new(string), new(string), new(string), new(string), },
			Path:                 "/apple/banana/cherry/grape",
			ExpectedArgs: []string{"apple","banana","cherry","grape"},
		},



		{
			Pattern: "/user/{sessionKey}",
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},
		{
			Pattern: "/user/{sessionKey}/",
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},

		{
			Pattern: "/user/{sessionKey}/vehicle",
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},
		{
			Pattern: "/user/{sessionKey}/vehicle/",
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},

		{
			Pattern: "/user/{sessionKey}/vehicle/DEFAULT",
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},
		{
			Pattern: "/user/{sessionKey}/vehicle/DEFAULT/",
			Args: []interface{}{new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij"},
		},

		{
			Pattern: "/user/{sessionKey}/vehicle/{vehicleIdcode}",
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "DEFAULT"},
		},
		{
			Pattern: "/user/{sessionKey}/vehicle/{vehicleIdcode}/",
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "DEFAULT"},
		},

		{
			Pattern: "/user/{sessionKey}/vehicle/{vehicleIdcode}",
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/N9Z_tiv7",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "N9Z_tiv7"},
		},
		{
			Pattern: "/user/{sessionKey}/vehicle/{vehicleIdcode}/",
			Args: []interface{}{new(string), new(string), },
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/N9Z_tiv7/",
			ExpectedArgs: []string{"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij", "N9Z_tiv7"},
		},



		{
			Pattern: "/-/object/{uuid}.jsonld",
			Args: []interface{}{new(string), },
			Path:                 "/-/object/ED7BA470-8E54-465E-825C-99712043E01C.jsonld",
			ExpectedArgs: []string{"ED7BA470-8E54-465E-825C-99712043E01C"},
		},



		{
			Pattern: "/-/outbox({begin},{end}).jsonld",
			Args: []interface{}{new(string), new(string), },
			Path:        "/-/outbox(12,345).jsonld",
			ExpectedArgs: []string{"12","345"},
		},
	}

	for testNumber, test := range tests {

		var pattern *pathmatch.Pattern = pathmatch.MustCompile(test.Pattern)

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

		if didMatch, err := pattern.Find(test.Path, test.Args...); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: %v", testNumber, err)
			continue
		} else if !didMatch {
			t.Errorf("For test #%d, expected path to match pattern, but it didn't.", testNumber)
			t.Logf("PATTERN: %q", test.Pattern)
			t.Logf("PATH:    %q", test.Path)
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
