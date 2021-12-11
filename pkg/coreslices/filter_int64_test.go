package coreslices

import "testing"

func TestInt64Filter(t *testing.T) {
	tables := []struct {
		slice  []int64
		filter func(int, int64) bool
		result []int64
	}{
		{
			slice:  []int64{0, 1, 2},
			filter:  func(i int, e int64) bool { return e >= 1 },
			result: []int64{1, 2},
		},
}

	for _, table := range tables {
		result := Int64Filter(table.slice, table.filter)
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
