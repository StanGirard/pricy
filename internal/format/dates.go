package format

import "fmt"

type DateInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
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
