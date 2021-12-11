package coreslices

import "testing"

func TestUint8In(t *testing.T) {
	tables := []struct{
		a      uint8
		slice  []uint8
		result bool
	}{
		{1, []uint8{1, 2, 3}, true},
		{7, []uint8{8, 9, 0}, false},
		{0, []uint8{}, false},
		{0, []uint8{1, 2}, false},
}

	for _, table := range tables {
		result := Uint8In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
