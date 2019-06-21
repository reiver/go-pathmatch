package pathmatch

import (
	"testing"
)

func TestInternalPatternAsPattern(t *testing.T) {

	var datum Pattern = new(internalPattern) // THIS IS WHAT ACTUALLY MATTERS.


	if nil == datum {
		t.Errorf("This should never happen.")
		return
	}
}
