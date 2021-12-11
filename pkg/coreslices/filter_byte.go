package coreslices

// ByteFilter godoc
//
// Filter slice of types byte.
func ByteFilter(slice []byte, filter func(int, byte) bool) []byte {
	result := make([]byte, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
