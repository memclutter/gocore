package coreslices

import "testing"

func TestStringApply(t *testing.T) {
	tables := []struct {
		slice  []string
		apply  func(int, string) string
		result []string
	}{
		{
			slice:  []string{"a", "b", "c"},
			apply:  func(i int, e string) string { return e + "1" },
			result: []string{"a1", "b1", "c1"},
		},
}

	for _, table := range tables {
		result := StringApply(table.slice, table.apply)
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
