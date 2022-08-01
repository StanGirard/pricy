package azure

import (
	"context"
	"flag"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/helpers"
)

type costUsageResult struct {
	cost         float64
	date         time.Time
	ResourceType string
	Currency     string
}

var (
	subscription = flag.String("subscription", "", "Subscription ID")
)

func Execute(configuration format.Configuration) []format.Service {
	flag.Parse()
	if *subscription == "" {
		panic("subscriptions is required with flag --subscriptions")
	}
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	// Set Subscription ID

	// Get context.Context
	if err != nil {
		panic(err)
	}

	// Generating Date
	var dateInterval format.DateInterval
	if configuration.Interval == "" {
		dateInterval = helpers.DaysInterval(configuration.Days)
	} else {
		var err error
		dateInterval, err = helpers.ParseInterval(configuration.Interval)
		if err != nil {
			panic(err)
		}
	}
	start, _ := time.Parse("2006-01-02", dateInterval.Start)
	end, _ := time.Parse("2006-01-02", dateInterval.End)

	costQueryClient, err := armcostmanagement.NewQueryClient(cred, nil)
	if err != nil {
		panic(err)
	}
	aggregations := make(map[string]*armcostmanagement.QueryAggregation)
	sum := armcostmanagement.QueryAggregation{
		Function: (*armcostmanagement.FunctionType)(to.StringPtr("Sum")),
		Name:     to.StringPtr("Cost"),
	}
	aggregations["totalCost"] = &sum

	var grouping []*armcostmanagement.QueryGrouping
	grouping = append(grouping, &armcostmanagement.QueryGrouping{
		Name: to.StringPtr("ResourceType"),
		Type: (*armcostmanagement.QueryColumnType)(to.StringPtr("Dimension")),
	})
	queryDefinition := armcostmanagement.QueryDefinition{
		Type:      (*armcostmanagement.ExportType)(to.StringPtr("ActualCost")),
		Timeframe: (*armcostmanagement.TimeframeType)(to.StringPtr("Custom")),
		Dataset: &armcostmanagement.QueryDataset{
			Granularity: (*armcostmanagement.GranularityType)(to.StringPtr(configuration.Granularity)),
			Aggregation: aggregations,
			Grouping:    grouping,
		},
		TimePeriod: &armcostmanagement.QueryTimePeriod{
			From: &start,
			To:   &end,
		},
	}
	subscriptionId := "/subscriptions/" + *subscription
	result, err := costQueryClient.Usage(context.Background(), subscriptionId, queryDefinition, nil)

	if err != nil {
		panic(err)
	}
	var costUsageResults []costUsageResult

	for _, v := range result.Properties.Rows {
		var result costUsageResult
		for i, v2 := range v {
			if i == 0 {
				cast, _ := v2.(float64)
				result.cost = cast
			} else if i == 1 {
				if str, ok := v2.(string); ok {
					// 2022-07-01T00:00:00
					date, _ := time.Parse("2006-01-02T15:04:05", str)
					result.date = date
				} else if str, ok := v2.(float64); ok {
					var y int64 = int64(str)
					timestring := strconv.Itoa(int(y))
					time, _ := time.Parse("20060102", timestring)
					result.date = time
				} else {
					panic("Unknown type")
				}

			}
			if i == 2 {
				cast, _ := v2.(string)
				result.ResourceType = cast
			}
			if i == 3 {
				cast, _ := v2.(string)
				result.Currency = cast
			}
		}
		costUsageResults = append(costUsageResults, result)
	}

	var services []format.Service
	for _, v := range costUsageResults {
		services = format.AddServices(services, v.ResourceType, v.Currency, format.DateInterval{Start: v.date.Format("2006-01-02"), End: v.date.Format("2006-01-02")}, v.cost)
	}
	return services
}
