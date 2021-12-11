package coreslices

import "testing"

func TestFloat64In(t *testing.T) {
	tables := []struct{
		a      float64
		slice  []float64
		result bool
	}{
		{.1, []float64{.1, .2, .3}, true},
		{.7, []float64{.8, .9}, false},
        {.01, []float64{}, false},
        {.01, []float64{.3, .2}, false},
}

	for _, table := range tables {
		result := Float64In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
