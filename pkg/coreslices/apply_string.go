package coreslices

// StringApply godoc
//
// Apply function for slice of types string.
func StringApply(slice []string, apply func(int, string)string) []string {
	result := make([]string, len(slice))
	for i, e := range slice {
		result[i] = apply(i, e)
	}
	return result
}
