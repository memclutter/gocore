package coreslices

import "testing"

func TestInt64Apply(t *testing.T) {
	tables := []struct {
		slice  []int64
		apply  func(int, int64) int64
		result []int64
	}{
		{
			slice:  []int64{0, 1, 2},
			apply:  func(i int, e int64) int64 { return e + 1 },
			result: []int64{1, 2, 3},
		},
}

	for _, table := range tables {
		result := Int64Apply(table.slice, table.apply)
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
