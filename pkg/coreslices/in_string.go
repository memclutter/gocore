package coreslices

// StringIn godoc
//
// Check some string tpl slice of string types exists.
func StringIn(a string, slice []string) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
