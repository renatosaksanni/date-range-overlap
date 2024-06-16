package daterange

import "time"

// DateRange represents a range of dates with a start and end date
type DateRange struct {
	Start time.Time
	End   time.Time
}

// Overlaps checks if two DateRanges overlap
func (d DateRange) Overlaps(other DateRange) bool {
	return d.End.After(other.Start) && d.Start.Before(other.End)
}
