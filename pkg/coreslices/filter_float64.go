package coreslices

// Float64Filter godoc
//
// Filter slice of types float64.
func Float64Filter(slice []float64, filter func(int, float64) bool) []float64 {
	result := make([]float64, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
