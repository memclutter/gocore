package coreslices

// StringFilter godoc
//
// Filter slice of types string.
func StringFilter(slice []string, filter func(int, string) bool) []string {
	result := make([]string, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
}
