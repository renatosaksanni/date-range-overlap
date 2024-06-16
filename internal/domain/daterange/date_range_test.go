package daterange

import (
	"testing"

	"github.com/renatosaksanni/date-range-overlap/pkg/utils"
)

func TestOverlaps(t *testing.T) {
	tests := []struct {
		name     string
		period1  DateRange
		period2  DateRange
		expected bool
	}{
		{"No overlap", DateRange{utils.ParseDate("2023-06-01"), utils.ParseDate("2023-06-10")}, DateRange{utils.ParseDate("2023-06-11"), utils.ParseDate("2023-06-20")}, false},
		{"Overlap", DateRange{utils.ParseDate("2023-06-01"), utils.ParseDate("2023-06-10")}, DateRange{utils.ParseDate("2023-06-05"), utils.ParseDate("2023-06-15")}, true},
		{"Touching but not overlapping", DateRange{utils.ParseDate("2023-06-01"), utils.ParseDate("2023-06-10")}, DateRange{utils.ParseDate("2023-06-10"), utils.ParseDate("2023-06-20")}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.period1.Overlaps(test.period2)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
