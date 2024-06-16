package utils

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	dateStr := "2023-06-01"
	expectedDate, _ := time.Parse("2006-01-02", dateStr)
	parsedDate := ParseDate(dateStr)

	if !parsedDate.Equal(expectedDate) {
		t.Fatalf("expected %v, got %v", expectedDate, parsedDate)
	}
}
