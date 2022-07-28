package aws

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"

	"github.com/stangirard/pricy/internal/dates"
	"github.com/stangirard/pricy/internal/format"
)

var (
	date  = flag.String("date", "", "date to get cost usage")
	month = flag.Bool("month", false, "get cost usage for month")
)

func createCostExplorer(sess *session.Session) *costexplorer.CostExplorer {
	return costexplorer.New(sess)
}

func getCostUsage(costExplorer *costexplorer.CostExplorer, start, end string) (*costexplorer.GetCostAndUsageOutput, error) {
	return costExplorer.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Granularity: aws.String("DAILY"),
		Metrics: []*string{
			aws.String("UNBLENDED_COST"),
		},
	})
}

func getCostUsageByService(costExplorer *costexplorer.CostExplorer, start, end string) (*costexplorer.GetCostAndUsageOutput, error) {
	return costExplorer.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Granularity: aws.String("MONTHLY"),
		Metrics: []*string{
			aws.String("UNBLENDED_COST"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
		},
	})
}

type ServicesPrices []format.ServicePrice

func parseCostUsagebyService(cost *costexplorer.GetCostAndUsageOutput, err error) []format.ServicePrice {
	var servicePrices []format.ServicePrice
	if err != nil {
		fmt.Println(err)
	}
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

func TotalCostUsage(costOutput *costexplorer.GetCostAndUsageOutput, err error) float64 {
	parsedCost := parseCostUsagebyService(costOutput, err)
	var totalCost float64
	for _, price := range parsedCost {
		totalCost = totalCost + price.Cost
	}
	return totalCost

}

func InitCostExplorer(session *session.Session) {
	// Initialize the session
	costExplorer := createCostExplorer(session)
	// Generating Date
	var dateInterval format.DateInterval
	if *month {
		dateInterval = dates.FromLastMonthToNow()
	} else {
		dateInterval = dates.FromLastWeekToNow()
	}
	costOutput, err := getCostUsageByService(costExplorer, dateInterval.Start, dateInterval.End)
	fmt.Println("Total Cost Usage: ", TotalCostUsage(costOutput, err), "From:", dateInterval.Start, "To:", dateInterval.End)
}
