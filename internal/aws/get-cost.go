package aws

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"

	"github.com/stangirard/pricy/internal/format"
)

var (
	date  = flag.String("date", "", "date to get cost usage")
	month = flag.Bool("month", false, "get cost usage for month")
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
	})
	if err != nil {
		panic(err)
	}
	return costAndUsageOutput

}
