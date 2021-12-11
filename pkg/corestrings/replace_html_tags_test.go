package corestrings

import "testing"

func TestReplaceHtmlTags(t *testing.T) {
	tables := []struct{
		s string
		r string
		result string
	}{
		{`<a href="http://example.com">Example.com</a>`, "", "Example.com"},
		{`<div><span>abcd</span><br />SomeText</div>`, "", "abcdSomeText"},
	}

	for _, table := range tables {
		result := ReplaceHtmlTags(table.s, table.r)
		if result != table.result {
			t.Errorf("replace tags on '%s' in string '%s' incorrect, except '%s' actual '%s'", table.r, table.s, table.result, result)
		}
	}
}
