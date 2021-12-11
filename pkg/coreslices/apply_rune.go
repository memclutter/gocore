package coreslices

// RuneApply godoc
//
// Apply function for slice of types rune.
func RuneApply(slice []rune, apply func(int, rune)rune) []rune {
	result := make([]rune, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
