package coreslices

// RuneIn godoc
//
// Check some rune tpl slice of rune types exists.
func RuneIn(a rune, slice []rune) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
