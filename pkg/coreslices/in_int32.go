package coreslices

// Int32In godoc
//
// Check some int32 tpl slice of int32 types exists.
func Int32In(a int32, slice []int32) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
