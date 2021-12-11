package coreslices

import "testing"

func TestUint64Filter(t *testing.T) {
	tables := []struct {
		slice  []uint64
		filter func(int, uint64) bool
		result []uint64
	}{
		{
			slice:  []uint64{0, 1, 2},
			filter:  func(i int, e uint64) bool { return e >= 1 },
			result: []uint64{1, 2},
		},
}

	for _, table := range tables {
		result := Uint64Filter(table.slice, table.filter)
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
