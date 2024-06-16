package utils

import "time"

// ParseDate parses a date string and returns a time.Time object
func ParseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(err)
	}
	return date
}
