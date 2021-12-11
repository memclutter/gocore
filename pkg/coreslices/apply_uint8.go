package coreslices

// Uint8Apply godoc
//
// Apply function for slice of types uint8.
func Uint8Apply(slice []uint8, apply func(int, uint8)uint8) []uint8 {
	result := make([]uint8, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
