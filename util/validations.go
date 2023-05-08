package util

import (
	"regexp"
	"strings"
)

// validate through REGEX if a string is a valid email address
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// checks if a string has at least 8 characters, has 1 either number or special char
func ValidatePassword(password string) bool {
	if len(strings.TrimSpace(password)) == 0 {
		return false
	}

	if len(password) < 8 {
		return false
	}

	// At least one special character or number
	pattern := "[^a-zA-Z]"
	matched, _ := regexp.MatchString(pattern, password)
	return matched
}
