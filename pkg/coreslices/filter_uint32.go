package coreslices

// Uint32Filter godoc
//
// Filter slice of types uint32.
func Uint32Filter(slice []uint32, filter func(int, uint32) bool) []uint32 {
	result := make([]uint32, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
