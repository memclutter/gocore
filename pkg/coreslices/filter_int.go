package coreslices

// IntFilter godoc
//
// Filter slice of types int.
func IntFilter(slice []int, filter func(int, int) bool) []int {
	result := make([]int, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
