package coreslices

import "testing"

func TestUint64Apply(t *testing.T) {
	tables := []struct {
		slice  []uint64
		apply  func(int, uint64) uint64
		result []uint64
	}{
		{
			slice:  []uint64{0, 1, 2},
			apply:  func(i int, e uint64) uint64 { return e + 1 },
			result: []uint64{1, 2, 3},
		},
}

	for _, table := range tables {
		result := Uint64Apply(table.slice, table.apply)
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
