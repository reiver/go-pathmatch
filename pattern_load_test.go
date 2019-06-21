package pathmatch


import (
	"github.com/fatih/structs"

	"testing"
)


func TestFindAndLoad(t *testing.T) {

	tests := []struct{
		Pattern            *Pattern
		StructPtr           interface{}
		Path                string
		Expected map[string]string
	}{
		{
			Pattern: MustCompile("/{this}/{that}/{these}/{those}"),
			StructPtr: new(struct{
				This  string `match:"this"`
				That  string `match:"that"`
				These string `match:"these"`
				Those string `match:"those"`
			}),
			Path:                 "/apple/banana/cherry/grape",
			Expected: map[string]string{
				"This":"apple",
				"That":"banana",
				"These":"cherry",
				"Those":"grape",
			},
		},



		{
			Pattern: MustCompile("/user/{sessionKey}"),
			StructPtr: new(struct{
				SessionKey string `match:"sessionKey"`
			}),
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			Expected: map[string]string{
				"SessionKey":"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/"),
			StructPtr: new(struct{
				SessionKey string `match:"sessionKey"`
			}),
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/",
			Expected: map[string]string{
				"SessionKey":"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			},
		},

		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle"),
			StructPtr: new(struct{
				SessionKey string `match:"sessionKey"`
			}),
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle",
			Expected: map[string]string{
				"SessionKey":"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/"),
			StructPtr: new(struct{
				SessionKey string `match:"sessionKey"`
			}),
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/",
			Expected: map[string]string{
				"SessionKey":"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			},
		},

		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/DEFAULT"),
			StructPtr: new(struct{
				SessionKey string `match:"sessionKey"`
			}),
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT",
			Expected: map[string]string{
				"SessionKey":"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			},
		},
		{
			Pattern: MustCompile("/user/{sessionKey}/vehicle/DEFAULT/"),
			StructPtr: new(struct{
				SessionKey string `match:"sessionKey"`
			}),
			Path:                 "/user/76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij/vehicle/DEFAULT/",
			Expected: map[string]string{
				"SessionKey":"76M6.mXQfgiGSC_YJ5uXSnWUmELbe8OgOm5n.iZ98Ij",
			},
		},
/*

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
*/
	}


	for testNumber, test := range tests {


		for structFieldName, structFieldValue := range structs.Map(test.StructPtr) {
			structFieldValueString, ok := structFieldValue.(string)
			if !ok {
				t.Errorf("For test #%d, expected test.StructPtr.%s to be of type string, but actually was %T.", testNumber, structFieldName, structFieldValue)
				continue
			}

			if expected, actual := "", structFieldValueString; expected != actual {
				t.Errorf("For test #%d, expected test.StructPtr.%s to (initially) be %q, but actually was %q.", testNumber, structFieldName, expected, actual)
				continue
			}
		}

		if didMatch, err := test.Pattern.FindAndLoad(test.Path, test.StructPtr); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: %v", testNumber, err)
			continue
		} else if !didMatch {
			t.Errorf("For test #%d, expected path to match pattern, but it didn't.", testNumber)
			continue
		}

		for structFieldName, structFieldValue := range structs.Map(test.StructPtr) {
			structFieldValueString, ok := structFieldValue.(string)
			if !ok {
				t.Errorf("For test #%d, expected test.StructPtr.%s to be of type string, but actually was %T.", testNumber, structFieldName, structFieldValue)
				continue
			}

			if expected, actual := test.Expected[structFieldName], structFieldValueString; expected != actual {
				t.Errorf("For test #%d, expected test.StructPtr.%s to (initially) be %q, but actually was %q.", testNumber, structFieldName, expected, actual)
				continue
			}
		}
	}
}
