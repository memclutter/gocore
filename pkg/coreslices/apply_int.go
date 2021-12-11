package coreslices

// IntApply godoc
//
// Apply function for slice of types int.
func IntApply(slice []int, apply func(int, int)int) []int {
	result := make([]int, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
