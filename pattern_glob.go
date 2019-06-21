package pathmatch

import (
	"bytes"
)

func (pattern *Pattern) Glob() string {
	if nil == pattern {
		return ""
	}

	pattern.mutex.RLock()
	defer pattern.mutex.RUnlock()

//@TODO: This shouldn't be executed every time!

	var buffer bytes.Buffer

	for _, bit := range pattern.bits {
		if wildcardBit == bit {
			buffer.WriteRune('*')
		} else {
			buffer.WriteString(bit)
		}
	}

	return buffer.String()
}
