package main

import (
	"fmt"

	"github.com/renatosaksanni/date-range-overlap/internal/domain/daterange"
	"github.com/renatosaksanni/date-range-overlap/pkg/utils"
)

func main() {
	period1 := daterange.DateRange{
		Start: utils.ParseDate("2023-06-01"),
		End:   utils.ParseDate("2023-06-10"),
	}
	period2 := daterange.DateRange{
		Start: utils.ParseDate("2023-06-05"),
		End:   utils.ParseDate("2023-06-15"),
	}

	if period1.Overlaps(period2) {
		fmt.Println("Period 1 overlaps with Period 2")
	} else {
		fmt.Println("Period 1 does not overlap with Period 2")
	}
}
