package pathmatch


import (
	"database/sql"
	"fmt"
)


func set(value string, argsIndex int, args ...interface{}) error {

	if 0 > argsIndex {
		return newInternalErrorComplainer("Index value %d is less than zero.", argsIndex)
	}

	if lenArgs := len(args); argsIndex >= lenArgs {
		expectedAtLeast := 1+argsIndex
		actual          := lenArgs

		return newNotEnoughArgumentsComplainer(expectedAtLeast, actual)
	}

	arg := args[argsIndex]

	switch a := arg.(type) {
	case sql.Scanner:
		if err := a.Scan(value); nil != err {
			return newScanError(err, argsIndex, fmt.Sprintf("%T", arg))
		}
	case *string:
		*a = value
	default:
		return newUnsupportedIndexedArgumentTypeComplainer(argsIndex, fmt.Sprintf("%T", arg))
	}

	return nil
}
