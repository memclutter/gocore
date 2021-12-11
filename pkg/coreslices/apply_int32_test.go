package coreslices

import "testing"

func TestInt32Apply(t *testing.T) {
	tables := []struct {
		slice  []int32
		apply  func(int, int32) int32
		result []int32
	}{
		{
			slice:  []int32{0, 1, 2},
			apply:  func(i int, e int32) int32 { return e + 1 },
			result: []int32{1, 2, 3},
		},
}

	for _, table := range tables {
		result := Int32Apply(table.slice, table.apply)
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
