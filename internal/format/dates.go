package format

import (
	"fmt"
	"sort"
)

func (dates DateIntervals) Headers() []string {
	var headers []string

	headers = append(headers, "Service")

	var datesHeader []string
	for _, date := range dates {
		datesHeader = append(datesHeader, date.Start)
	}

	sort.Slice(datesHeader, func(i, j int) bool {
		return datesHeader[i] < datesHeader[j]
	})

	headers = append(headers, datesHeader...)
	return headers
}

func (d DateInterval) GetStartDate() string {
	return d.Start
}

func (d DateInterval) GetEndDate() string {
	return d.End
}

func (d DateInterval) Print() {
	fmt.Printf("%s - %s\n", d.Start, d.End)
}

func (d DateInterval) String() string {
	return fmt.Sprintf("%s - %s", d.Start, d.End)
}

func ContainsDateInterval(dates []DateInterval, date DateInterval) bool {
	for _, d := range dates {
		if d.Start == date.Start && d.End == date.End {
			return true
		}
	}
	return false
}

func SortDates(dates []DateInterval) []DateInterval {
	for i := 0; i < len(dates); i++ {
		for j := i + 1; j < len(dates); j++ {
			if dates[i].Start > dates[j].Start {
				dates[i], dates[j] = dates[j], dates[i]
			}
		}
	}
	return dates
}
