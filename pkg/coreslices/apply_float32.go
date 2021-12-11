package coreslices

// Float32Apply godoc
//
// Apply function for slice of types float32.
func Float32Apply(slice []float32, apply func(int, float32)float32) []float32 {
	result := make([]float32, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
