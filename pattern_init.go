package pathmatch

func (receiver *Pattern) init(fieldTagName string) {
	if nil == receiver {
		return
	}

	bits     := []string{}
	names    := []string{}
	namesSet := map[string]struct{}{}

	receiver.bits         = bits
	receiver.names        = names
	receiver.namesSet     = namesSet
	receiver.fieldTagName = fieldTagName
}
