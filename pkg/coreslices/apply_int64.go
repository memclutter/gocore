package coreslices

// Int64Apply godoc
//
// Apply function for slice of types int64.
func Int64Apply(slice []int64, apply func(int, int64)int64) []int64 {
	result := make([]int64, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
