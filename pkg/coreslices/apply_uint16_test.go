package coreslices

import "testing"

func TestUint16Apply(t *testing.T) {
	tables := []struct {
		slice  []uint16
		apply  func(int, uint16) uint16
		result []uint16
	}{
		{
			slice:  []uint16{0, 1, 2},
			apply:  func(i int, e uint16) uint16 { return e + 1 },
			result: []uint16{1, 2, 3},
		},
}

	for _, table := range tables {
		result := Uint16Apply(table.slice, table.apply)
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
