package coreslices

// Uint8Filter godoc
//
// Filter slice of types uint8.
func Uint8Filter(slice []uint8, filter func(int, uint8) bool) []uint8 {
	result := make([]uint8, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
