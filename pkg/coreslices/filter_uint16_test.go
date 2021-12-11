package coreslices

import "testing"

func TestUint16Filter(t *testing.T) {
	tables := []struct {
		slice  []uint16
		filter func(int, uint16) bool
		result []uint16
	}{
		{
			slice:  []uint16{0, 1, 2},
			filter:  func(i int, e uint16) bool { return e >= 1 },
			result: []uint16{1, 2},
		},
}

	for _, table := range tables {
		result := Uint16Filter(table.slice, table.filter)
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
