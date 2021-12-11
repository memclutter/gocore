package coreslices

// Uint8In godoc
//
// Check some uint8 tpl slice of uint8 types exists.
func Uint8In(a uint8, slice []uint8) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
