package aws

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stangirard/pricy/internal/format"
)

func getCostUsageByService(costExplorer *costexplorer.CostExplorer, start, end string, granularity string) *costexplorer.GetCostAndUsageOutput {
	costAndUsageOutput, err := costExplorer.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Granularity: aws.String(granularity),
		Metrics: []*string{
			aws.String("UNBLENDED_COST"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
		},
		Filter: &costexplorer.Expression{
			Not: &costexplorer.Expression{
				Dimensions: &costexplorer.DimensionValues{
					Key:    aws.String("RECORD_TYPE"),
					Values: []*string{aws.String("REFUND"), aws.String("Credit"), aws.String("SavingsPlanNegation"), aws.String("Tax")},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return costAndUsageOutput

}

func FormatCostUsagebyService(cost *costexplorer.GetCostAndUsageOutput) []format.Service {
	var services []format.Service
	for _, result := range cost.ResultsByTime {
		for _, group := range result.Groups {
			for index, metric := range group.Metrics {
				if index == "UnblendedCost" {
					cost, _ := strconv.ParseFloat(*metric.Amount, 64)
					services = format.AddServices(services, *group.Keys[0], *metric.Unit, format.DateInterval{Start: *result.TimePeriod.Start, End: *result.TimePeriod.End}, cost)
				}
			}
		}
	}
	return services
}
