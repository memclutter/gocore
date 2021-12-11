package coreslices

import "testing"

func TestFloat32Apply(t *testing.T) {
	tables := []struct {
		slice  []float32
		apply  func(int, float32) float32
		result []float32
	}{
		{
			slice:  []float32{0.00, 0.01, 0.02},
			apply:  func(i int, e float32) float32 { return e + 0.01 },
			result: []float32{0.01, 0.02, 0.03},
		},
}

	for _, table := range tables {
		result := Float32Apply(table.slice, table.apply)
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
