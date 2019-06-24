package pathmatch

// String makes pathmatch.Pattern fit the fmt.Stringer interface.
//
// String returns the (pre-compiled) pattern template.
func (receiver Pattern) String() string {
	return receiver.template
}
