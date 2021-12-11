package coreslices

// Int64In godoc
//
// Check some int64 tpl slice of int64 types exists.
func Int64In(a int64, slice []int64) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
