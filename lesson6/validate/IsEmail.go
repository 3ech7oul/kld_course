package validate

import "regexp"

// Check email by regEx.
// In addition, the email addresses have a practical limit of 254 bytes.
func IsEmail(s string) bool {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(s) > 254 || !rxEmail.MatchString(s) {
		return false
	}
	return true
}
