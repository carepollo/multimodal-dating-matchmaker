package util

import "strconv"

// function that checks if number can be converted into integer
func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
