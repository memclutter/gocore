package coreslices

import "testing"

func TestUint16In(t *testing.T) {
	tables := []struct{
		a      uint16
		slice  []uint16
		result bool
	}{
		{1, []uint16{1, 2, 3}, true},
		{7, []uint16{8, 9, 0}, false},
		{0, []uint16{}, false},
		{0, []uint16{1, 2}, false},
}

	for _, table := range tables {
		result := Uint16In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
