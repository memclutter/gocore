package coreslices

// Uint64Apply godoc
//
// Apply function for slice of types uint64.
func Uint64Apply(slice []uint64, apply func(int, uint64)uint64) []uint64 {
	result := make([]uint64, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
