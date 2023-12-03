package helpers

import "regexp"

// validate through REGEX if a string is a valid email address
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	return matched && err == nil
}

// checks if a string has at least 8 characters at least 1 number and at least one uppercase
func ValidatePassword(password string) bool {
	pattern := `^[A-Za-z0-9]*[A-Z]+[A-Za-z0-9]*[0-9]+[A-Za-z0-9]*$`
	matched, err := regexp.MatchString(pattern, password)
	return matched && err == nil
}
