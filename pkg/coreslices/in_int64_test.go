package coreslices

import "testing"

func TestInt64In(t *testing.T) {
	tables := []struct{
		a      int64
		slice  []int64
		result bool
	}{
		{1, []int64{1, 2, 3}, true},
		{7, []int64{8, 9, 0}, false},
		{0, []int64{}, false},
		{0, []int64{-1, -2}, false},
}

	for _, table := range tables {
		result := Int64In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
