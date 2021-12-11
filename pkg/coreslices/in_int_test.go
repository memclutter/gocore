package coreslices

import "testing"

func TestIntIn(t *testing.T) {
	tables := []struct{
		a      int
		slice  []int
		result bool
	}{
		{1, []int{1, 2, 3}, true},
		{7, []int{8, 9, 0}, false},
		{0, []int{}, false},
		{0, []int{-1, -2}, false},
}

	for _, table := range tables {
		result := IntIn(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
