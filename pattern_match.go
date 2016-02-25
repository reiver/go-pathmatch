package pathmatch


import (
	"strings"
)


const (
	doesNotMatter = false
)


var (
	errThisShouldNeverHappen = newInternalErrorComplainer("This should never happen.")
)


func (pattern *internalPattern) Match(path string, args ...interface{}) (bool, error) {

	s := path

	argsIndex := 0
	for _, bit := range pattern.bits {

		switch bit {
		default:
			if !strings.HasPrefix(s, bit) {
				return false, nil
			}

			s = s[len(bit):]
		case wildcardBit:
			index := strings.IndexRune(s, '/')
			if -1 == index {
				if err := set(s, argsIndex, args...); nil != err {
					return doesNotMatter, err
				}
				argsIndex++
			} else if 0 <= index {
				value := s[:index]
				if err := set(value, argsIndex, args...); nil != err {
					return doesNotMatter, err
				}
				argsIndex++
				s = s[index:]
			} else {
				return doesNotMatter, errThisShouldNeverHappen
			}
		}

	}


	return true, nil
}
