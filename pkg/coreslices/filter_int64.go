package coreslices

// Int64Filter godoc
//
// Filter slice of types int64.
func Int64Filter(slice []int64, filter func(int, int64) bool) []int64 {
	result := make([]int64, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
