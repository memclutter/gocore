package corestrings

import "testing"

func TestCamelToSnake(t *testing.T) {
	tables := []struct {
		camel  string
		result string
	}{
		{"CamelCase", "camel_case"},
		{"ApiKey", "api_key"},
	}

	for _, table := range tables {
		result := CamelToSnake(table.camel)
		if result != table.result {
			t.Errorf("incorrect convert string '%s' CamelCase -> snake_case, excepted '%s', actual '%s'", table.camel, table.result, result)
		}
	}
}
