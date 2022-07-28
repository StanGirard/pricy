package aws

import (
	"flag"
	"fmt"
	"sort"

	"github.com/stangirard/pricy/internal/format"
)

var (
	details = flag.Bool("details", false, "print details")
)

func PrintCostByService(PricePerDate []format.PricePerDate) {
	//Sort by DateInterval Start
	TotalCostUsage := TotalCostUsage(PricePerDate)
	sort.Slice(PricePerDate, func(i, j int) bool {
		return PricePerDate[i].DateInterval.Start < PricePerDate[j].DateInterval.Start
	})

	for _, pricesPerDay := range PricePerDate {
		//Print in Bold
		fmt.Printf("%s - %s: Price is %f\n", pricesPerDay.DateInterval.Start, pricesPerDay.DateInterval.End, FindPriceForDateInterval(TotalCostUsage, pricesPerDay.DateInterval))
		fmt.Println("-----------------------------")

		sort.Slice(pricesPerDay.ServicePrice, func(i, j int) bool {
			return pricesPerDay.ServicePrice[i].Cost > pricesPerDay.ServicePrice[j].Cost
		})
		if *details {
			for _, servicePrice := range pricesPerDay.ServicePrice {
				//Sort by cost

				fmt.Printf("%s: %f%s\n", servicePrice.Service, servicePrice.Cost, format.UnitsToSymbol[servicePrice.Units])
			}
		}
	}

}

func PrintTotalCost(totalCost float64) {
	fmt.Println("Total Cost:", totalCost)
}

func reportGenerate(costByServices []format.PricePerDate) {

	PrintCostByService(costByServices)

	// PrintTotalCost(TotalCostUsage(costByServices))
}
