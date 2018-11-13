// Package leap implements a utility to determine leap years
package leap

// IsLeapYear takes a year as input and returns true if it is a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if (year%4 == 0) && (year%100 != 0) {
		return true
	}
	return false
}
