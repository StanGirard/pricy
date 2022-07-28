package aws

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stangirard/pricy/internal/format"
)

func formatCostUsagebyService(cost *costexplorer.GetCostAndUsageOutput) []format.ServicePrice {
	var servicePrices []format.ServicePrice

	for _, costDetail := range cost.ResultsByTime {
		for _, group := range costDetail.Groups {
			// Print UNBLENDED_COST and BLENDED_COST with name and value

			var servicePrice format.ServicePrice
			servicePrice.Service = *group.Keys[0]
			for index, metric := range group.Metrics {
				if index == "UnblendedCost" {
					cost, _ := strconv.ParseFloat(*metric.Amount, 64)
					servicePrice.Cost = cost
					servicePrice.Units = *metric.Unit
				}
			}
			servicePrices = append(servicePrices, servicePrice)

		}
	}
	return servicePrices
}

func TotalCostUsage(ServicesPrices []format.ServicePrice) float64 {
	var totalCost float64
	for _, price := range ServicesPrices {
		totalCost = totalCost + price.Cost
	}
	return totalCost

}
