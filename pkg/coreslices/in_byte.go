package coreslices

// ByteIn godoc
//
// Check some byte tpl slice of byte types exists.
func ByteIn(a byte, slice []byte) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
