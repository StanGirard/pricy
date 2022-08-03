package reports

import (
	"github.com/stangirard/pricy/internal/format"
)

// Calculates the evolution of the prices by interval
func (services Services) calculateEvolution() {
	dates := format.FindDatesIntervals(services)
	dates = format.SortDates(dates)
	for _, service := range services {
		for index, date := range dates {
			if index == 0 {
				continue
			}
			previousDate := dates[index-1]
			if service.DatePrice[previousDate] == 0.0 {
				if service.DatePrice[date] == 0.0 {
					service.PriceEvolution[date] = 0.0
					continue
				} else {
					service.PriceEvolution[date] = 999.0
					continue
				}
			} else if service.DatePrice[date] == 0.0 {
				if service.DatePrice[previousDate] != 0.0 {
					service.PriceEvolution[date] = -999.0
					continue
				}
			} else {
				service.PriceEvolution[date] = (service.DatePrice[date] - service.DatePrice[previousDate]) / service.DatePrice[previousDate]
			}
		}
	}
}
