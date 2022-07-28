package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"

	"github.com/stangirard/pricy/internal/format"
)

type ServicesPrices []format.ServicePrice

func getCostUsageByService(costExplorer *costexplorer.CostExplorer, start, end string) *costexplorer.GetCostAndUsageOutput {
	costAndUsageOutput, err := costExplorer.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
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
		Filter: &costexplorer.Expression{
			Not: &costexplorer.Expression{
				Dimensions: &costexplorer.DimensionValues{
					Key:    aws.String("RECORD_TYPE"),
					Values: []*string{aws.String("REFUND"), aws.String("Credit"), aws.String("SavingsPlanNegation")},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return costAndUsageOutput

}
