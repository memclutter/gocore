package coreslices

// Uint16Filter godoc
//
// Filter slice of types uint16.
func Uint16Filter(slice []uint16, filter func(int, uint16) bool) []uint16 {
	result := make([]uint16, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
