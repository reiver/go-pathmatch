package pathmatch

// Match returns true if ‘path’ matches the compiled pattern, else returns false if it doesn't match.
func (receiver *Pattern) Match(path string) (bool, error) {
	if nil == receiver {
		return false, errNilReceiver
	}

//@TODO: Is it a good idea to be dynamically creating this?
//@TODO: Also, can the struct fields be put in here directly instead?
	args := []interface{}{}
	numNames := len(receiver.MatchNames())
	for i:=0; i<numNames; i++ {
		args = append(args, new(string))
	}

	found, err := receiver.Find(path, args...)
	if nil != err {
		return false, err
	}

	return found, nil
}
