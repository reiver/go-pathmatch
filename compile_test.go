package pathmatch


import (
	"testing"
)


func TestCompileAndMatchNames(t *testing.T) {

	tests := []struct{
		UncompiledPattern           string
		ExpectedNames             []string
		ExpectedNamesSet map[string]struct{}
		ExpectedBits              []string
	}{
		{
			UncompiledPattern: "",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"",
			},
		},



		{
			UncompiledPattern: "/",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"/",
			},
		},



		{
			UncompiledPattern: "//",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"//",
			},
		},



		{
			UncompiledPattern: "///",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"///",
			},
		},



		{
			UncompiledPattern:   "/{this}",
			ExpectedNames:[]string{"this"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "/{this}/",
			ExpectedNames:[]string{"this"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
				"/",
			},
		},
		{
			UncompiledPattern:   "/{this}/{that}",
			ExpectedNames:[]string{"this","that"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
				"/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "/{this}/{that}/",
			ExpectedNames:[]string{"this","that"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
			},
		},
		{
			UncompiledPattern:   "/{this}/{that}/{these}",
			ExpectedNames:[]string{"this","that","these"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "/{this}/{that}/{these}/",
			ExpectedNames:[]string{"this","that","these"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
			},
		},
		{
			UncompiledPattern:   "/{this}/{that}/{these}/{those}",
			ExpectedNames:[]string{"this","that","these","those"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
				"those":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "/{this}/{that}/{these}/{those}/",
			ExpectedNames:[]string{"this","that","these","those"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
				"those":struct{}{},
			},
			ExpectedBits: []string{
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
			},
		},



		{
			UncompiledPattern:   "{this}",
			ExpectedNames:[]string{"this"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "{this}/",
			ExpectedNames:[]string{"this"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
				"/",
			},
		},
		{
			UncompiledPattern:   "{this}/{that}",
			ExpectedNames:[]string{"this","that"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
				"/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "{this}/{that}/",
			ExpectedNames:[]string{"this","that"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
				"/",
				wildcardBit,
				"/",
			},
		},
		{
			UncompiledPattern:   "{this}/{that}/{these}",
			ExpectedNames:[]string{"this","that","these"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "{this}/{that}/{these}/",
			ExpectedNames:[]string{"this","that","these"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
			},
		},
		{
			UncompiledPattern:   "{this}/{that}/{these}/{those}",
			ExpectedNames:[]string{"this","that","these","those"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
				"those":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern:   "{this}/{that}/{these}/{those}/",
			ExpectedNames:[]string{"this","that","these","those"},
			ExpectedNamesSet:map[string]struct{}{
				"this":struct{}{},
				"that":struct{}{},
				"these":struct{}{},
				"those":struct{}{},
			},
			ExpectedBits: []string{
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
				wildcardBit,
				"/",
			},
		},



		{
			UncompiledPattern: "/apple/banana/cherry",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"/apple/banana/cherry",
			},
		},
		{
			UncompiledPattern: "/apple/banana/cherry/",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"/apple/banana/cherry/",
			},
		},
		{
			UncompiledPattern: "//apple/banana/cherry",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"//apple/banana/cherry",
			},
		},
		{
			UncompiledPattern: "/apple//banana/cherry",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"/apple//banana/cherry",
			},
		},
		{
			UncompiledPattern: "/apple/banana/cherry//",
			ExpectedNames:[]string{},
			ExpectedNamesSet:map[string]struct{}{},
			ExpectedBits: []string{
				"/apple/banana/cherry//",
			},
		},



		{
			UncompiledPattern: "/apple/{banana}",
			ExpectedNames:[]string{    "banana"},
			ExpectedNamesSet:map[string]struct{}{
				"banana":struct{}{},
			},
			ExpectedBits: []string{
				"/apple/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern: "/apple/{banana}/",
			ExpectedNames:[]string{    "banana"},
			ExpectedNamesSet:map[string]struct{}{
				"banana":struct{}{},
			},
			ExpectedBits: []string{
				"/apple/",
				wildcardBit,
				"/",
			},
		},
		{
			UncompiledPattern: "/apple/{banana}//",
			ExpectedNames:[]string{    "banana"},
			ExpectedNamesSet:map[string]struct{}{
				"banana":struct{}{},
			},
			ExpectedBits: []string{
				"/apple/",
				wildcardBit,
				"//",
			},
		},



		{
			UncompiledPattern: "/apple/{banana}/cherry",
			ExpectedNames:[]string{    "banana"},
			ExpectedNamesSet:map[string]struct{}{
				"banana":struct{}{},
			},
			ExpectedBits: []string{
				"/apple/",
				wildcardBit,
				"/cherry",
			},
		},
		{
			UncompiledPattern: "/apple/{banana}/cherry/",
			ExpectedNames:[]string{    "banana"},
			ExpectedNamesSet:map[string]struct{}{
				"banana":struct{}{},
			},
			ExpectedBits: []string{
				"/apple/",
				wildcardBit,
				"/cherry/",
			},
		},



		{
			UncompiledPattern: "/apple/{banana}/cherry/{grape}",
			ExpectedNames:[]string{    "banana",       "grape"},
			ExpectedNamesSet:map[string]struct{}{
				"banana":struct{}{},
				"grape":struct{}{},
			},
			ExpectedBits: []string{
				"/apple/",
				wildcardBit,
				"/cherry/",
				wildcardBit,
			},
		},
		{
			UncompiledPattern: "/apple/{banana}/cherry/{grape}/",
			ExpectedNames:[]string{    "banana",       "grape"},
			ExpectedNamesSet:map[string]struct{}{
				"banana":struct{}{},
				"grape":struct{}{},
			},
			ExpectedBits: []string{
				"/apple/",
				wildcardBit,
				"/cherry/",
				wildcardBit,
				"/",
			},
		},
	}


	for testNumber, test := range tests {

		var actualPattern Pattern

		err := CompileTo(&actualPattern, test.UncompiledPattern)
		if nil != err {
			t.Errorf("For test #%d, did not expect to receive an error, but actually got one: %v\nPATTERN: %q", testNumber, err, test.UncompiledPattern)
			continue
		}

		if expected, actual := test.ExpectedBits, actualPattern.bits; len(expected) != len(actual) {
			t.Errorf("For test #%d, expected compiled pattern to have %d bits, but actually had %d.\nEXPECTED BITS: %#v\nACTUAL BITS:   %#v\nPATTERN: %q", testNumber, len(expected), len(actual), expected, actual, test.UncompiledPattern)
			continue
		} else {
			for bitNumber, expectedBit := range expected {

				actualBit := actual[bitNumber]

				if expectedBit != actualBit {
					t.Errorf("For test #%d and bit #%d, expected %q, but actually got %q.\nEXPECTED BITS: %#v\nACTUAL BITS:   %#v\nPATTERN: %q", testNumber, bitNumber, expectedBit, actualBit, expected, actual, test.UncompiledPattern)
					continue
				}
			}
		}

		if expected, actual := test.ExpectedNames, actualPattern.MatchNames(); len(expected) != len(actual) {
			t.Errorf("For test #%d, when checking directly, expected compiled pattern to have %d names, but actually had %d.\nEXPECTED NAMES: %#v\nACTUAL NAMES:   %#v\nPATTERN: %q", testNumber, len(expected), len(actual), expected, actual, test.UncompiledPattern)
			continue
		} else {
			for nameNumber, expectedName := range expected {

				actualName := actual[nameNumber]

				if expectedName != actualName {
					t.Errorf("For test #%d and name #%d, expected %q, but actually got %q.\nEXPECTED NAMES: %#v\nACTUAL NAMES:   %#v\nPATTERN: %q", testNumber, nameNumber, expectedName, actualName, expected, actual, test.UncompiledPattern)
					continue
				}
			}
		}

		if expected, actual := test.ExpectedNamesSet, actualPattern.namesSet; len(expected) != len(actual) {
			t.Errorf("For test #%d, when checking directly, expected compiled pattern to have %d names in set, but actually had %d.\nEXPECTED NAMES SET: %#v\nACTUAL NAMES SET:   %#v\nPATTERN: %q", testNumber, len(expected), len(actual), expected, actual, test.UncompiledPattern)
			continue
		} else {
			for expectedName, _ := range expected {

				_, ok := actual[expectedName]

				if !ok {
					t.Errorf("For test #%d, expected name %q to exist, but actually did't.\nEXPECTED NAMES SET: %#v\nACTUAL NAMES SET:   %#v\nPATTERN: %q", testNumber, expectedName, expected, actual, test.UncompiledPattern)
					continue
				}
			}
		}

		if expected, actual := len(test.ExpectedNames), len(actualPattern.MatchNames()); expected != actual {
			t.Errorf("For test #%d, when checking using .MatchNames(), expected compiled pattern to have %d names, but actually had %d.\nEXPECTED NAMES: %#v\nACTUAL NAMES:   %#v\nPATTERN: %q", testNumber, expected, actual, test.ExpectedNames, actualPattern.MatchNames(), test.UncompiledPattern)
			continue
		}
	}
}


func TestCompileFail(t *testing.T) {

	tests := []struct{
		UncompiledPattern string
		ExpectedError     string
	}{
		{
			UncompiledPattern: "/users/{userId",
			ExpectedError:     `Bad Request: Syntax Error: Missing ending "}" (to match beginning "{").`,
		},
		{
			UncompiledPattern: "/users/{userId}/vehicles/{vehicleId",
			ExpectedError:     `Bad Request: Syntax Error: Missing ending "}" (to match beginning "{").`,
		},
		{
			UncompiledPattern: "/users/{userId/vehicles/{vehicleId}",
			ExpectedError:     `Bad Request: Syntax Error: "/" inside of "{...}".`,
		},
		{
			UncompiledPattern: "/users/{userId{vehicleId}",
			ExpectedError:     `Bad Request: Syntax Error: "{" inside of "{...}".`,
		},
		{
			UncompiledPattern: "/apple/{fruitId}/banana/{fruitId}",
			ExpectedError:     `Bad Request: Syntax Error: Duplicate match name: "fruitId".`,
		},
		{
			UncompiledPattern: "/fruit/{apple/banana/cherry}",
			ExpectedError:     `Bad Request: Syntax Error: "/" inside of "{...}".`,
		},
	}


	for testNumber, test := range tests {
		var pattern Pattern

		err := CompileTo(&pattern, test.UncompiledPattern)
		if nil == err {
			t.Errorf("For test #%d, expected to receive an error, but actually did not get one: %v\nPATTERN: %q", testNumber, err, test.UncompiledPattern)
			continue
		}
		if expected, actual := test.ExpectedError, err.Error(); expected != actual {
			t.Errorf("For test #%d, expected error message to be %q, but actually was %q.\nPATTERN: %q", testNumber, expected, actual, test.UncompiledPattern)
			continue
		}
	}
}
