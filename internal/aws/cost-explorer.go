package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
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
		Granularity: aws.String("MONTHLY"),
		Metrics: []*string{
			aws.String("BLENDED_COST"),
			aws.String("UNBLENDED_COST"),
			aws.String("AMOUNT_REFUNDED"),
			aws.String("NET_RISK"),
		},
	})
}

type DateInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func InitCostExplorer() *costexplorer.GetCostAndUsageOutput {
	sess := createSession()
	costExplorer := createCostExplorer(sess)
	cost, err := getCostUsage(costExplorer, "2022-07-01", "2022-07-20")
	if err != nil {
		panic(err)
	}
	return cost
}
