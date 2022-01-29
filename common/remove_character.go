package common

import (
	"strings"
	"unicode"
)

// RemoveCharacterInString remove character in string
func RemoveCharacterInString(s string, substr string) string {
	if n := strings.Index(s, substr); n >= 0 {
		return strings.TrimRightFunc(s[:n], unicode.IsSpace)
	}
	return s
}
