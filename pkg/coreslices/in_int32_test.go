package coreslices

import "testing"

func TestInt32In(t *testing.T) {
	tables := []struct{
		a      int32
		slice  []int32
		result bool
	}{
		{1, []int32{1, 2, 3}, true},
		{7, []int32{8, 9, 0}, false},
		{0, []int32{}, false},
		{0, []int32{-1, -2}, false},
}

	for _, table := range tables {
		result := Int32In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
