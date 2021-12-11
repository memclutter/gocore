package coreslices

// Int32Filter godoc
//
// Filter slice of types int32.
func Int32Filter(slice []int32, filter func(int, int32) bool) []int32 {
	result := make([]int32, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
