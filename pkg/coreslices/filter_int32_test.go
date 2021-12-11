package coreslices

import "testing"

func TestInt32Filter(t *testing.T) {
	tables := []struct {
		slice  []int32
		filter func(int, int32) bool
		result []int32
	}{
		{
			slice:  []int32{0, 1, 2},
			filter:  func(i int, e int32) bool { return e >= 1 },
			result: []int32{1, 2},
		},
}

	for _, table := range tables {
		result := Int32Filter(table.slice, table.filter)
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
