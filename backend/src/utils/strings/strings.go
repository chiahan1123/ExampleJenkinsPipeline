package utils

import (
	"strconv"
	"strings"
)

// NumberToString converts a floating number to string with given number of digits,
// the string contains no digits if the number is exactly an integer.
func NumberToString(number float64, digits int) string {
	if digits < 0 {
		return ""
	}
	if number == float64(int(number)) {
		return strconv.FormatFloat(number, 'f', -1, 64)
	}
	return strings.TrimRight(strconv.FormatFloat(number, 'f', digits, 64), "0")
}
