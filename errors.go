package pathmatch

import (
	"errors"
)

var (
	errNilReceiver = errors.New("pathmatch: Nil Receiver")
	errNilTarget   = errors.New("pathmatch: Nil Target")
)
