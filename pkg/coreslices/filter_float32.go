package coreslices

// Float32Filter godoc
//
// Filter slice of types float32.
func Float32Filter(slice []float32, filter func(int, float32) bool) []float32 {
	result := make([]float32, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
