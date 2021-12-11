package coreslices

import "testing"

func TestStringFilter(t *testing.T) {
	tables := []struct {
		slice  []string
		filter func(int, string) bool
		result []string
	}{
		{
			slice:  []string{"a", "b", "c"},
			filter:  func(i int, e string) bool { return e != "b" },
			result: []string{"a", "c"},
		},
}

	for _, table := range tables {
		result := StringFilter(table.slice, table.filter)
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
