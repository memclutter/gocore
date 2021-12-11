package coreslices

// Uint32In godoc
//
// Check some uint32 tpl slice of uint32 types exists.
func Uint32In(a uint32, slice []uint32) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
