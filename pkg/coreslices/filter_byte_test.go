package coreslices

import "testing"

func TestByteFilter(t *testing.T) {
	tables := []struct {
		slice  []byte
		filter func(int, byte) bool
		result []byte
	}{
		{
			slice:  []byte{0x00, 0x01, 0x02},
			filter:  func(i int, e byte) bool { return e >= 0x01 },
			result: []byte{0x01, 0x02},
		},
}

	for _, table := range tables {
		result := ByteFilter(table.slice, table.filter)
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
