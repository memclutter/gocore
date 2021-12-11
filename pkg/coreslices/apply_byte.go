package coreslices

// ByteApply godoc
//
// Apply function for slice of types byte.
func ByteApply(slice []byte, apply func(int, byte)byte) []byte {
	result := make([]byte, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
