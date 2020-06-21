// Package validate provides string validators
package validate

import "unicode"

// Convert string to runes and check all runge in Latin symbol
func IsString(s string) bool {
	valid := true
	for _, r := range []rune(s) {
		if !unicode.Is(unicode.Latin, r) {
			valid = false
		}
	}
	return valid
}
