package reports

import (
	"flag"
	"fmt"
	"sort"

	"github.com/stangirard/pricy/internal/format"
)

var (
	details = flag.Bool("details", false, "print details") // Print more information in your terminal
)

// Name of a service and price for a given date (not specified)
type priceToDate struct {
	Name  string
	Price float64
}

// Array of pricetodate
type priceToDateArray []priceToDate

// Array of services
type Services []format.Service

// Print in the terminal the cost of each services for each day intervals
func (Services Services) printCost() {
	allDates := format.SortDates(format.FindDatesIntervals(Services))
	for _, date := range allDates {

		var pricesByDate priceToDateArray
		for _, service := range Services {
			pricesByDate = append(pricesByDate, priceToDate{Name: service.Name, Price: service.DatePrice[date]})
		}
		sort.Slice(pricesByDate, func(i, j int) bool {
			return pricesByDate[i].Price > pricesByDate[j].Price
		})

		totalCost := pricesByDate.getCostByDate()
		fmt.Printf("%s: %.2f%s \n", date.String(), totalCost, format.UnitsToSymbol[Services[0].Units])
		if *details {
			for _, service := range pricesByDate {
				fmt.Printf("\t%s: %.2f%s\n", service.Name, service.Price, format.UnitsToSymbol[Services[0].Units])
			}
		}

	}

}

// Get the cost of a service for a given date
func (costByDate priceToDateArray) getCostByDate() float64 {
	var total float64
	for _, service := range costByDate {
		total += service.Price
	}
	// Round to 2 decimals
	return float64(int(total*100)) / 100
}

// Initialise the report module
func InitReport(services Services) {
	flag.Parse()
	services.calculateEvolution()
	services.printCost()
	services.initCSV()
	services.initHTML()
}
