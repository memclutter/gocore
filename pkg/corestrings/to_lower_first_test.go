package corestrings

import "testing"

func TestToLowerFirst(t *testing.T)  {
	tables := []struct{
		s string
		result string
	}{
		{
			s: "SomeStringValue",
			result: "someStringValue",
		},
		{
			s: "",
			result: "",
		},
		{
			s: "aBar",
			result: "aBar",
		},
	}

	for _, table := range tables {
		result := ToLowerFirst(table.s)
		if table.result != result {
			t.Errorf("lower first char incorrect, excepted '%s', actual '%s'", table.result, result)
		}
	}
}
