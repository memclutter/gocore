package coreslices

import "testing"

func TestUint64In(t *testing.T) {
	tables := []struct{
		a      uint64
		slice  []uint64
		result bool
	}{
		{1, []uint64{1, 2, 3}, true},
		{7, []uint64{8, 9, 0}, false},
		{0, []uint64{}, false},
		{0, []uint64{1, 2}, false},
}

	for _, table := range tables {
		result := Uint64In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
