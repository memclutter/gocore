package coreslices

// Uint16Apply godoc
//
// Apply function for slice of types uint16.
func Uint16Apply(slice []uint16, apply func(int, uint16)uint16) []uint16 {
	result := make([]uint16, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
