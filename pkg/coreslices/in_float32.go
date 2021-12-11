package coreslices

// Float32In godoc
//
// Check some float32 tpl slice of float32 types exists.
func Float32In(a float32, slice []float32) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
