package format

import (
	"fmt"
	"sort"
)

// Creates the headers for the CSV export. It appends Service at the beginning and then gets all the dates.
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

// Get the start date of a date interval
func (d DateInterval) GetStartDate() string {
	return d.Start
}

// Get the end date of a date interval
func (d DateInterval) GetEndDate() string {
	return d.End
}

// Prints the date interval
func (d DateInterval) Print() {
	fmt.Printf("%s - %s\n", d.Start, d.End)
}

// Turns the date interval into a string
func (d DateInterval) String() string {
	return fmt.Sprintf("%s - %s", d.Start, d.End)
}

// Check wether an array of date interval contains a dateinterval
func ContainsDateInterval(dates []DateInterval, date DateInterval) bool {
	for _, d := range dates {
		if d.Start == date.Start && d.End == date.End {
			return true
		}
	}
	return false
}

// Sort dates in a date interval array
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
