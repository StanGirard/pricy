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

type PriceToDate struct {
	Name  string
	Price float64
}

func PrintCostByService(Services []format.Service) {
	allDates := format.SortDates(format.FindAllDateIntervals(Services))
	for _, date := range allDates {

		var serviceToDate []PriceToDate
		for _, service := range Services {
			serviceToDate = append(serviceToDate, PriceToDate{Name: service.Service, Price: service.DatePrice[date]})
		}
		sort.Slice(serviceToDate, func(i, j int) bool {
			return serviceToDate[i].Price > serviceToDate[j].Price
		})

		totalCost := GetCostByDateAllServices(serviceToDate)
		fmt.Printf("%s: %.2f%s \n", date.String(), totalCost, format.UnitsToSymbol[Services[0].Units])
		if *details {
			for _, service := range serviceToDate {
				fmt.Printf("\t%s: %.2f%s\n", service.Name, service.Price, format.UnitsToSymbol[Services[0].Units])
			}
		}

	}

}

func GetCostByDateAllServices(costByDate []PriceToDate) float64 {
	var total float64
	for _, service := range costByDate {
		total += service.Price
	}
	// Round to 2 decimals
	return float64(int(total*100)) / 100
}

func GenerateReport(costByServices []format.Service) {
	flag.Parse()
	GenerateEvolutionFromPreviousDate(costByServices)
	PrintCostByService(costByServices)
	GenerateCSVReport(costByServices)
	GenerateEvolutionFromPreviousDate(costByServices)

}
