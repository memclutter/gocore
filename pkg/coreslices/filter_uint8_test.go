package coreslices

import "testing"

func TestUint8Filter(t *testing.T) {
	tables := []struct {
		slice  []uint8
		filter func(int, uint8) bool
		result []uint8
	}{
		{
			slice:  []uint8{0, 1, 2},
			filter:  func(i int, e uint8) bool { return e >= 1 },
			result: []uint8{1, 2},
		},
}

	for _, table := range tables {
		result := Uint8Filter(table.slice, table.filter)
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
