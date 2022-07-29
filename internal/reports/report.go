package reports

import (
	"flag"
	"fmt"
	"sort"

	"github.com/stangirard/pricy/internal/format"
)

var (
	details = flag.Bool("details", false, "print details")
)

type priceToDate struct {
	Name  string
	Price float64
}

type priceToDateArray []priceToDate

type ServicesArray []format.Service

func (Services ServicesArray) printCost() {
	allDates := format.SortDates(format.FindDatesIntervals(Services))
	for _, date := range allDates {

		var pricesByDate priceToDateArray
		for _, service := range Services {
			pricesByDate = append(pricesByDate, priceToDate{Name: service.Service, Price: service.DatePrice[date]})
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

func (costByDate priceToDateArray) getCostByDate() float64 {
	var total float64
	for _, service := range costByDate {
		total += service.Price
	}
	// Round to 2 decimals
	return float64(int(total*100)) / 100
}

func InitReport(services ServicesArray) {
	flag.Parse()
	services.calculateEvolution()
	services.printCost()
	services.initCSV()
}
