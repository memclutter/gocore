package coreslices

import "testing"

func TestIntApply(t *testing.T) {
	tables := []struct {
		slice  []int
		apply  func(int, int) int
		result []int
	}{
		{
			slice:  []int{0, 1, 2},
			apply:  func(i int, e int) int { return e + 1 },
			result: []int{1, 2, 3},
		},
}

	for _, table := range tables {
		result := IntApply(table.slice, table.apply)
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
