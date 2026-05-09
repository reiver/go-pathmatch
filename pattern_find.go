package pathmatch

import (
	"strings"
	"unicode/utf8"
)

const (
	doesNotMatter = false
)

var (
	errThisShouldNeverHappen = newInternalError("This should never happen.")
)

// Find compares ‘path’ against its (compiled) pattern template; if it matches it loads the
// matches into ‘args’, and then returns true.
//
// Find may set some, or all of the items in ‘args’ even if it returns false, and even if it
// returns an error.
func (pattern *Pattern) Find(path string, args ...interface{}) (bool, error) {
	if nil == pattern {
		return false, errNilReceiver
	}

	pattern.mutex.RLock()
	defer pattern.mutex.RUnlock()

	s := path

	argsIndex := 0
	for bitIndex, bit := range pattern.bits {

		switch bit {
		default:
			if !strings.HasPrefix(s, bit) {
				return false, nil
			}

			s = s[len(bit):]
		case wildcardBit:
			if "" == s {
				return false, nil
			}

			index := strings.IndexRune(s, '/')
			{
				var nextBitIndex int = 1+bitIndex

				if nextBitIndex < len(pattern.bits) {
					nextBit := pattern.bits[nextBitIndex]

					r, _ := utf8.DecodeRuneInString(nextBit)

					index = strings.IndexRune(s, r)
				}
			}

			var value string
			switch {
			default:
				return doesNotMatter, errThisShouldNeverHappen
			case -1 == index:
				value = s
			case 0 <= index:
				value = s[:index]
			}

			if err := set(value, argsIndex, args...); nil != err {
				return doesNotMatter, err
			}
			argsIndex++
			s = s[len(value):]
		}
	}

	if "" != s {
		return false, nil
	}

	return true, nil
}
