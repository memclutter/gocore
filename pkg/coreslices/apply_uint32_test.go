package coreslices

import "testing"

func TestUint32Apply(t *testing.T) {
	tables := []struct {
		slice  []uint32
		apply  func(int, uint32) uint32
		result []uint32
	}{
		{
			slice:  []uint32{0, 1, 2},
			apply:  func(i int, e uint32) uint32 { return e + 1 },
			result: []uint32{1, 2, 3},
		},
}

	for _, table := range tables {
		result := Uint32Apply(table.slice, table.apply)
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
