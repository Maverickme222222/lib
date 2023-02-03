package shared

import "regexp"

// ValidAlphaNumeric checks validity of strings submitted
func ValidAlphaNumeric(value string) bool {
	var IsLetter = regexp.MustCompile(`^[-a-zA-Z0-9_]*$`).MatchString
	var IsNumber = regexp.MustCompile(`^[0-9]*$`).MatchString

	if !IsLetter(value) {
		return false
	}

	if IsNumber(value) {
		return false
	}

	return true
}
