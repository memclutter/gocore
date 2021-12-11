package coreslices

// Int32Apply godoc
//
// Apply function for slice of types int32.
func Int32Apply(slice []int32, apply func(int, int32)int32) []int32 {
	result := make([]int32, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
