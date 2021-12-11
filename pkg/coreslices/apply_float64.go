package coreslices

// Float64Apply godoc
//
// Apply function for slice of types float64.
func Float64Apply(slice []float64, apply func(int, float64)float64) []float64 {
	result := make([]float64, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
