package coreslices

// RuneFilter godoc
//
// Filter slice of types rune.
func RuneFilter(slice []rune, filter func(int, rune) bool) []rune {
	result := make([]rune, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
