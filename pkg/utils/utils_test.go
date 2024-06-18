package utils

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Time
	}{
		{"2023-06-01", time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)},
		{"2024-01-01", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		result := ParseDate(test.input)
		if !result.Equal(test.expected) {
			t.Errorf("ParseDate(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestParseDateInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ParseDate did not panic on invalid input")
		}
	}()
	ParseDate("invalid-date")
}
