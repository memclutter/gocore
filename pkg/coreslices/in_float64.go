package coreslices

// Float64In godoc
//
// Check some float64 tpl slice of float64 types exists.
func Float64In(a float64, slice []float64) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
