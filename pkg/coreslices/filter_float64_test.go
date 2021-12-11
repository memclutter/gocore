package coreslices

import "testing"

func TestFloat64Filter(t *testing.T) {
	tables := []struct {
		slice  []float64
		filter func(int, float64) bool
		result []float64
	}{
		{
			slice:  []float64{0.00, 0.01, 0.02},
			filter:  func(i int, e float64) bool { return e >= 0.01 },
			result: []float64{0.01, 0.02},
		},
}

	for _, table := range tables {
		result := Float64Filter(table.slice, table.filter)
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
