package coreslices

import "testing"

func TestRuneIn(t *testing.T) {
	tables := []struct{
		a      rune
		slice  []rune
		result bool
	}{
		{'a', []rune{'a', 'b', 'c'}, true},
		{'w', []rune{'x', 'y', 'z'}, false},
        {0x00, []rune{}, false},
        {0x00, []rune{0x00}, true},
}

	for _, table := range tables {
		result := RuneIn(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
