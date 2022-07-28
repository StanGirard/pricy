package aws

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stangirard/pricy/internal/format"
)

func formatCostUsagebyService(cost *costexplorer.GetCostAndUsageOutput) []format.PricePerDate {
	var pricePerDate []format.PricePerDate

	for _, result := range cost.ResultsByTime {
		var servicePrices []format.ServicePrice
		for _, group := range result.Groups {
			for index, metric := range group.Metrics {
				if index == "UnblendedCost" {
					cost, _ := strconv.ParseFloat(*metric.Amount, 64)
					servicePrices = append(servicePrices, format.ServicePrice{
						Service: *group.Keys[0],
						Cost:    cost,
						Units:   *metric.Unit,
					})
				}
			}
		}
		pricePerDate = append(pricePerDate, format.PricePerDate{
			DateInterval: format.DateInterval{
				Start: *result.TimePeriod.Start,
				End:   *result.TimePeriod.End,
			},
			ServicePrice: servicePrices,
		})
	}
	return pricePerDate
}

type TotalPerDay struct {
	DateInterval format.DateInterval
	TotalCost    float64
}

func TotalCostUsage(PricePerDate []format.PricePerDate) []TotalPerDay {
	var TotalPerDays []TotalPerDay
	for _, pricesPerDay := range PricePerDate {
		var totalCost float64
		for _, servicePrice := range pricesPerDay.ServicePrice {
			totalCost += servicePrice.Cost
		}
		TotalPerDays = append(TotalPerDays, TotalPerDay{
			DateInterval: pricesPerDay.DateInterval,
			TotalCost:    totalCost,
		})
	}
	return TotalPerDays
}

func FindPriceForDateInterval(PricePerDate []TotalPerDay, dateInterval format.DateInterval) float64 {
	for _, pricesPerDay := range PricePerDate {
		if pricesPerDay.DateInterval.Start == dateInterval.Start && pricesPerDay.DateInterval.End == dateInterval.End {
			return pricesPerDay.TotalCost
		}
	}
	return 0
}
