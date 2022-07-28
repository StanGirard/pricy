package format

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func FormatCostUsagebyService(cost *costexplorer.GetCostAndUsageOutput) []Service {
	var services []Service
	for _, result := range cost.ResultsByTime {
		for _, group := range result.Groups {
			for index, metric := range group.Metrics {
				if index == "UnblendedCost" {
					cost, _ := strconv.ParseFloat(*metric.Amount, 64)
					services = AddServices(services, *group.Keys[0], *metric.Unit, DateInterval{Start: *result.TimePeriod.Start, End: *result.TimePeriod.End}, cost)
				}
			}
		}
	}
	return services
}

type TotalPerDay struct {
	Date      DateInterval
	TotalCost float64
}

func TotalCostUsage(Services []Service) []TotalPerDay {
	var TotalPerDays []TotalPerDay
	for _, service := range Services {
		for date, price := range service.DatePrice {
			for i := range TotalPerDays {
				if TotalPerDays[i].Date.Start == date.Start && TotalPerDays[i].Date.End == date.End {
					TotalPerDays[i].TotalCost += price
					break
				} else {
					TotalPerDays = append(TotalPerDays, TotalPerDay{Date: date, TotalCost: price})
					break
				}
			}
		}
	}
	return TotalPerDays
}
