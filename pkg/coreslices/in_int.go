package coreslices

// IntIn godoc
//
// Check some int tpl slice of int types exists.
func IntIn(a int, slice []int) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
