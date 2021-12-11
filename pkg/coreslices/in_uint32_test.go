package coreslices

import "testing"

func TestUint32In(t *testing.T) {
	tables := []struct{
		a      uint32
		slice  []uint32
		result bool
	}{
		{1, []uint32{1, 2, 3}, true},
		{7, []uint32{8, 9, 0}, false},
		{0, []uint32{}, false},
		{0, []uint32{1, 2}, false},
}

	for _, table := range tables {
		result := Uint32In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
