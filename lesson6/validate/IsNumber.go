package validate

import "strconv"

// Convert string to numeric. If string have been converted success number this is number.
func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
