package coreslices

import "testing"

func TestByteApply(t *testing.T) {
	tables := []struct {
		slice  []byte
		apply  func(int, byte) byte
		result []byte
	}{
		{
			slice:  []byte{0x00, 0x01, 0x02},
			apply:  func(i int, e byte) byte { return e + 0x01 },
			result: []byte{0x01, 0x02, 0x03},
		},
}

	for _, table := range tables {
		result := ByteApply(table.slice, table.apply)
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
