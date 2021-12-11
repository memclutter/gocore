package coreslices

import "testing"

func TestByteIn(t *testing.T) {
	tables := []struct{
		a      byte
		slice  []byte
		result bool
	}{
		{0x01, []byte{0x01, 0x02, 0x03}, true},
		{0x01, []byte{0x08, 0x09, 0x00}, false},
        {0x00, []byte{}, false},
        {0x00, []byte{0x00}, true},
}

	for _, table := range tables {
		result := ByteIn(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
