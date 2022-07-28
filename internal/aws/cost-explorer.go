package aws

import (
	"fmt"

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
			aws.String("BLENDED_COST"),
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

type ServicePrice struct {
	Service        string
	BLENDED_COST   string
	UNBLENDED_COST string
}

type ServicesPrices []ServicePrice

// define print function for ServicePrice
func (s ServicePrice) print() {
	fmt.Printf("Service: %s, BLENDED_COST: %s, UNBLENDED_COST: %s\n", s.Service, s.BLENDED_COST, s.UNBLENDED_COST)
}

func parseCostUsagebyService(cost *costexplorer.GetCostAndUsageOutput, err error) []ServicePrice {
	var servicePrices []ServicePrice
	if err != nil {
		fmt.Println(err)
	}
	for _, costDetail := range cost.ResultsByTime {
		for _, group := range costDetail.Groups {
			// Print UNBLENDED_COST and BLENDED_COST with name and value

			var servicePrice ServicePrice
			servicePrice.Service = *group.Keys[0]
			for index, metric := range group.Metrics {
				if index == "BlendedCost" {
					servicePrice.BLENDED_COST = *metric.Amount
				} else if index == "UnblendedCost" {
					servicePrice.UNBLENDED_COST = *metric.Amount
				}
			}
			servicePrices = append(servicePrices, servicePrice)

		}
	}
	return servicePrices
}

type DateInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func InitCostExplorer(session *session.Session) {
	costExplorer := createCostExplorer(session)
	prices := parseCostUsagebyService(getCostUsageByService(costExplorer, "2022-07-01", "2022-07-20"))
	for _, price := range prices {
		price.print()
	}
}
