package coreslices

// Uint64Filter godoc
//
// Filter slice of types uint64.
func Uint64Filter(slice []uint64, filter func(int, uint64) bool) []uint64 {
	result := make([]uint64, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
