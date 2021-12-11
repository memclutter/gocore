package coreslices

import "testing"

func TestIntFilter(t *testing.T) {
	tables := []struct {
		slice  []int
		filter func(int, int) bool
		result []int
	}{
		{
			slice:  []int{0, 1, 2},
			filter:  func(i int, e int) bool { return e >= 1 },
			result: []int{1, 2},
		},
}

	for _, table := range tables {
		result := IntFilter(table.slice, table.filter)
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
