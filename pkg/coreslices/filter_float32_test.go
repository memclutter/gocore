package coreslices

import "testing"

func TestFloat32Filter(t *testing.T) {
	tables := []struct {
		slice  []float32
		filter func(int, float32) bool
		result []float32
	}{
		{
			slice:  []float32{0.00, 0.01, 0.02},
			filter:  func(i int, e float32) bool { return e >= 0.01 },
			result: []float32{0.01, 0.02},
		},
}

	for _, table := range tables {
		result := Float32Filter(table.slice, table.filter)
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
