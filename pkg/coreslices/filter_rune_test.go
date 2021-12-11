package coreslices

import "testing"

func TestRuneFilter(t *testing.T) {
	tables := []struct {
		slice  []rune
		filter func(int, rune) bool
		result []rune
	}{
		{
			slice:  []rune{0x31, 0x32, 0x33},
			filter:  func(i int, e rune) bool { return e >= 0x32 },
			result: []rune{0x32, 0x33},
		},
}

	for _, table := range tables {
		result := RuneFilter(table.slice, table.filter)
		if len(table.result) != len(result) {
			t.Fatalf("excepted %d elements in result, but %d elements actual", len(table.result), len(result))
		}
		for i, e := range result {
			ee := table.result[i]
			if e != ee {
				t.Errorf("excepted %d element %v, but %v element actual", i, ee, e)
			}
		}
	}
}
