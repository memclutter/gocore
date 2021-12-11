package coreslices

import "testing"

func TestUint32Filter(t *testing.T) {
	tables := []struct {
		slice  []uint32
		filter func(int, uint32) bool
		result []uint32
	}{
		{
			slice:  []uint32{0, 1, 2},
			filter:  func(i int, e uint32) bool { return e >= 1 },
			result: []uint32{1, 2},
		},
}

	for _, table := range tables {
		result := Uint32Filter(table.slice, table.filter)
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
