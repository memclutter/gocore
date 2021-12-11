package corestrings

import (
	"strings"
	"unicode"
)

// ToLowerFirst godoc
//
// Make first char of string to lower
func ToLowerFirst(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	for i, ch := range s {
		return string(unicode.ToLower(ch)) + s[i+1:]
	}
	return s
}
