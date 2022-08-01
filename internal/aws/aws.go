package aws

import (
	"flag"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/helpers"
)

func createCostExplorer(sess *session.Session) *costexplorer.CostExplorer {
	return costexplorer.New(sess)
}

func Execute(configuration format.Configuration) []format.Service {
	// Initialize the session
	flag.Parse()
	session := initSession()
	costExplorer := createCostExplorer(session)

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

	// uppercase string for the granularity
	costUsageByService := getCostUsageByService(costExplorer, dateInterval.Start, dateInterval.End, strings.ToUpper(configuration.Granularity))
	formatCostUsagebyService := formatCostUsagebyService(costUsageByService)
	return formatCostUsagebyService
}
