package coreslices

// Uint16In godoc
//
// Check some uint16 tpl slice of uint16 types exists.
func Uint16In(a uint16, slice []uint16) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
