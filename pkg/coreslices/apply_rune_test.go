package coreslices

import "testing"

func TestRuneApply(t *testing.T) {
	tables := []struct {
		slice  []rune
		apply  func(int, rune) rune
		result []rune
	}{
		{
			slice:  []rune{0x31, 0x32, 0x33},
			apply:  func(i int, e rune) rune { return e + 0x01 },
			result: []rune{0x32, 0x33, 0x34},
		},
}

	for _, table := range tables {
		result := RuneApply(table.slice, table.apply)
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
