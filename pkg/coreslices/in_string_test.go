package coreslices

import "testing"

func TestStringIn(t *testing.T) {
	tables := []struct{
		a      string
		slice  []string
		result bool
	}{
		{"a", []string{"a", "b", "c"}, true},
		{"w", []string{"x", "y", "z"}, false},
        {"empty", []string{}, false},
        {"", []string{"empty", "blank"}, false},
}

	for _, table := range tables {
		result := StringIn(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
