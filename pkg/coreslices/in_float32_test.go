package coreslices

import "testing"

func TestFloat32In(t *testing.T) {
	tables := []struct{
		a      float32
		slice  []float32
		result bool
	}{
		{.1, []float32{.1, .2, .3}, true},
		{.7, []float32{.8, .9}, false},
        {.01, []float32{}, false},
        {.01, []float32{.3, .2}, false},
}

	for _, table := range tables {
		result := Float32In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
