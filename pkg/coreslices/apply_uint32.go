package coreslices

// Uint32Apply godoc
//
// Apply function for slice of types uint32.
func Uint32Apply(slice []uint32, apply func(int, uint32)uint32) []uint32 {
	result := make([]uint32, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
