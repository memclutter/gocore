package coreslices

// Uint64In godoc
//
// Check some uint64 tpl slice of uint64 types exists.
func Uint64In(a uint64, slice []uint64) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
