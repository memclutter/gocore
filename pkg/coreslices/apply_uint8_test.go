package coreslices

import "testing"

func TestUint8Apply(t *testing.T) {
	tables := []struct {
		slice  []uint8
		apply  func(int, uint8) uint8
		result []uint8
	}{
		{
			slice:  []uint8{0, 1, 2},
			apply:  func(i int, e uint8) uint8 { return e + 1 },
			result: []uint8{1, 2, 3},
		},
}

	for _, table := range tables {
		result := Uint8Apply(table.slice, table.apply)
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
