package pathmatch


import (
	"math/rand"
	"time"

	"testing"
)


func TestSetStringPointers(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
	runes := []rune(".0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz")
	lenRunes := len(runes)

	randomString := func() string {
		length := 1 + randomness.Intn(59)
		rs := make([]rune, length)
		for i := range rs {
			rs[i] = runes[randomness.Intn(lenRunes)]
		}

		return string(rs)
	}


	tests := []struct{
		Value     string
		ArgsIndex int
		Args    []interface{}
	}{
		{
			Value: randomString(),
			ArgsIndex: 0,
			Args: []interface{}{ new(string), },
		},



		{
			Value: randomString(),
			ArgsIndex: 0,
			Args: []interface{}{ new(string), new(string), },
		},
		{
			Value: randomString(),
			ArgsIndex: 1,
			Args: []interface{}{ new(string), new(string), },
		},



		{
			Value: randomString(),
			ArgsIndex: 0,
			Args: []interface{}{ new(string), new(string), new(string), },
		},
		{
			Value: randomString(),
			ArgsIndex: 1,
			Args: []interface{}{ new(string), new(string), new(string), },
		},
		{
			Value: randomString(),
			ArgsIndex: 2,
			Args: []interface{}{ new(string), new(string), new(string), },
		},
	}


	for testNumber, test := range tests {

		for argNumber, arg := range test.Args {
			if argStringPtr, ok := arg.(*string); !ok {
				t.Errorf("For test #%d, expected test.Args[%d] to be type *string (pointer to string), but actually was: %T", testNumber, argNumber, arg)
				continue
			} else if argString := *argStringPtr; "" != argString {
				t.Errorf("For test #%d, expected test.Args[%d] to be \"\" (empty string), but actually was: %q", testNumber, argNumber, argString)
				continue
			}
		}

		if err := set(test.Value, test.ArgsIndex, test.Args...); nil != err {
			t.Errorf("For test #%d, did not expected error when calling set(), but actually received one: %v", testNumber, err)
			continue
		}

		for argNumber, arg := range test.Args {
			if argNumber == test.ArgsIndex {
				continue
			}

			if argStringPtr, ok := arg.(*string); !ok {
				t.Errorf("For test #%d, expected test.Args[%d] to be type *string (pointer to string), but actually was: %T", testNumber, argNumber, arg)
				continue
			} else if argString := *argStringPtr; "" != argString {
				t.Errorf("For test #%d, expected test.Args[%d] to be \"\" (empty string), but actually was: %q", testNumber, argNumber, argString)
				continue
			}
		}

		if actualPtr, ok := test.Args[test.ArgsIndex].(*string); !ok {
			t.Errorf("For test #%d, expected test.Args[%d] to be type *string (pointer to string), but actually was: %T", testNumber, test.ArgsIndex, test.Args[test.ArgsIndex])
			continue
		} else if expected, actual := test.Value, *actualPtr; expected != actual {
			t.Errorf("For test #%d, expected test.Args[%d] to be type %q, but actually was %q.", testNumber, test.ArgsIndex, expected, actual)
			continue
		}
	}
}


func TestSetFail(t *testing.T) {

	const internalError           = "internal_error"
	const notEnoughArguments      = "not_enough_arguments"
	const unsupportedArgumentType = "unsupported_argument_type"


	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
	runes := []rune(".0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz")
	lenRunes := len(runes)

	randomString := func() string {
		length := 1 + randomness.Intn(59)
		rs := make([]rune, length)
		for i := range rs {
			rs[i] = runes[randomness.Intn(lenRunes)]
		}

		return string(rs)
	}


	tests := []struct{
		Value               string
		ArgsIndex           int
		Args              []interface{}
		ExpectedFit         string
		ExpectedErrorString string
	}{
		{
			Value: randomString(),
			ArgsIndex: -1,
			Args: []interface{}{ new(string), },
			ExpectedFit: internalError,
			ExpectedErrorString: "Internal Error: Index value -1 is less than zero.",
		},
		{
			Value: randomString(),
			ArgsIndex: -1,
			Args: []interface{}{ randomString(), },
			ExpectedFit: internalError,
			ExpectedErrorString: "Internal Error: Index value -1 is less than zero.",
		},



		{
			Value: randomString(),
			ArgsIndex: -2,
			Args: []interface{}{ new(string), },
			ExpectedFit: internalError,
			ExpectedErrorString: "Internal Error: Index value -2 is less than zero.",
		},
		{
			Value: randomString(),
			ArgsIndex: -2,
			Args: []interface{}{ randomString(), },
			ExpectedFit: internalError,
			ExpectedErrorString: "Internal Error: Index value -2 is less than zero.",
		},



		{
			Value: randomString(),
			ArgsIndex: -3,
			Args: []interface{}{ new(string), },
			ExpectedFit: internalError,
			ExpectedErrorString: "Internal Error: Index value -3 is less than zero.",
		},
		{
			Value: randomString(),
			ArgsIndex: -3,
			Args: []interface{}{ randomString(), },
			ExpectedFit: internalError,
			ExpectedErrorString: "Internal Error: Index value -3 is less than zero.",
		},



		{
			Value: randomString(),
			ArgsIndex: 0,
			Args: []interface{}{ randomString(), },
			ExpectedFit: unsupportedArgumentType,
			ExpectedErrorString: "Bad Request: Type of argument #0 (string) is unsupported. However, type \"*string\" (pointer to string) is supported; did you mean to use a \"*string\" instead?",
		},



		{
			Value: randomString(),
			ArgsIndex: 0,
			Args: []interface{}{ randomString(), randomString(), },
			ExpectedFit: unsupportedArgumentType,
			ExpectedErrorString: "Bad Request: Type of argument #0 (string) is unsupported. However, type \"*string\" (pointer to string) is supported; did you mean to use a \"*string\" instead?",
		},
		{
			Value: randomString(),
			ArgsIndex: 1,
			Args: []interface{}{ randomString(), randomString(), },
			ExpectedFit: unsupportedArgumentType,
			ExpectedErrorString: "Bad Request: Type of argument #1 (string) is unsupported. However, type \"*string\" (pointer to string) is supported; did you mean to use a \"*string\" instead?",
		},



		{
			Value: randomString(),
			ArgsIndex: 0,
			Args: []interface{}{ randomString(), randomString(), randomString(), },
			ExpectedFit: unsupportedArgumentType,
			ExpectedErrorString: "Bad Request: Type of argument #0 (string) is unsupported. However, type \"*string\" (pointer to string) is supported; did you mean to use a \"*string\" instead?",
		},
		{
			Value: randomString(),
			ArgsIndex: 1,
			Args: []interface{}{ randomString(), randomString(), randomString(), },
			ExpectedFit: unsupportedArgumentType,
			ExpectedErrorString: "Bad Request: Type of argument #1 (string) is unsupported. However, type \"*string\" (pointer to string) is supported; did you mean to use a \"*string\" instead?",
		},
		{
			Value: randomString(),
			ArgsIndex: 2,
			Args: []interface{}{ randomString(), randomString(), randomString(), },
			ExpectedFit: unsupportedArgumentType,
			ExpectedErrorString: "Bad Request: Type of argument #2 (string) is unsupported. However, type \"*string\" (pointer to string) is supported; did you mean to use a \"*string\" instead?",
		},



		{
			Value: randomString(),
			ArgsIndex: 0,
			Args: []interface{}{ },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 1 argument, but actually got 0.",
		},
		{
			Value: randomString(),
			ArgsIndex: 1,
			Args: []interface{}{ },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 2 arguments, but actually got 0.",
		},
		{
			Value: randomString(),
			ArgsIndex: 2,
			Args: []interface{}{ },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 3 arguments, but actually got 0.",
		},



		{
			Value: randomString(),
			ArgsIndex: 1,
			Args: []interface{}{ randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 2 arguments, but actually got 1.",
		},
		{
			Value: randomString(),
			ArgsIndex: 2,
			Args: []interface{}{ randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 3 arguments, but actually got 1.",
		},
		{
			Value: randomString(),
			ArgsIndex: 3,
			Args: []interface{}{ randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 4 arguments, but actually got 1.",
		},



		{
			Value: randomString(),
			ArgsIndex: 2,
			Args: []interface{}{ randomString(), randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 3 arguments, but actually got 2.",
		},
		{
			Value: randomString(),
			ArgsIndex: 3,
			Args: []interface{}{ randomString(), randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 4 arguments, but actually got 2.",
		},
		{
			Value: randomString(),
			ArgsIndex: 4,
			Args: []interface{}{ randomString(), randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 5 arguments, but actually got 2.",
		},



		{
			Value: randomString(),
			ArgsIndex: 3,
			Args: []interface{}{ randomString(), randomString(), randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 4 arguments, but actually got 3.",
		},
		{
			Value: randomString(),
			ArgsIndex: 4,
			Args: []interface{}{ randomString(), randomString(), randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 5 arguments, but actually got 3.",
		},
		{
			Value: randomString(),
			ArgsIndex: 5,
			Args: []interface{}{ randomString(), randomString(), randomString(), },
			ExpectedFit: notEnoughArguments,
			ExpectedErrorString: "Bad Request: Not enough arguments. Expected at least 6 arguments, but actually got 3.",
		},
	}


	for testNumber, test := range tests {

		if err := set(test.Value, test.ArgsIndex, test.Args...); nil == err {
			t.Errorf("For test #%d, expected error when calling set(), but did not actually received one: %v", testNumber, err)
			continue
		} else {
			switch err.(type) {
			case InternalErrorComplainer:
				if expected, actual := internalError, test.ExpectedFit; expected != actual {
					t.Errorf("For test #%d, did indeed expect an error, but did not expect it to fit the \"InternalErrorComplainer\" interface, but actually did: %T.", testNumber, err)
					continue
				}
			case NotEnoughArgumentsComplainer:
				if expected, actual := notEnoughArguments, test.ExpectedFit; expected != actual {
					t.Errorf("For test #%d, did indeed expect an error, but did not expect it to fit the \"NotEnoughArgumentsComplainer\" interface, but actually did: %T.", testNumber, err)
					continue
				}
			case UnsupportedArgumentType:
				if expected, actual := unsupportedArgumentType, test.ExpectedFit; expected != actual {
					t.Errorf("For test #%d, did indeed expect an error, but did not expect it to fit the \"UnsupportedArgumentType\" interface, but actually did: %T.", testNumber, err)
					continue
				}
			default:
				t.Errorf("For test #%d, did indeed expect an error, but did not expect this error: [%v] (%T).", testNumber, err, err)
				continue
			}

			if expected, actual := test.ExpectedErrorString, err.Error(); expected != actual {
				t.Errorf("For test #%d, expected error string to be %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
		}

	}
}
