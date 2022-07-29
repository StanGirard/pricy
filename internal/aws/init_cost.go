package aws

import (
	"flag"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/helpers"
	"github.com/stangirard/pricy/internal/reports"
)

var (
	granularity = flag.String("granularity", "DAILY", "granularity to get cost usage")
	days        = flag.Int("days", 14, "get cost usage for last 14 days")
	interval    = flag.String("interval", "", "get cost usage for a specific interval as '2022-03-30:2022-03-31' ")
)

func createCostExplorer(sess *session.Session) *costexplorer.CostExplorer {
	return costexplorer.New(sess)
}

func InitCostExplorer(session *session.Session) {
	// Initialize the session
	flag.Parse()
	costExplorer := createCostExplorer(session)

	// Generating Date
	var dateInterval format.DateInterval
	if interval == nil || *interval == "" {
		dateInterval = helpers.DaysInterval(*days)
	} else {
		dateInterval = helpers.ParseInterval(*interval)
	}

	// uppercase string for the granularity
	costUsageByService := getCostUsageByService(costExplorer, dateInterval.Start, dateInterval.End, strings.ToUpper(*granularity))
	formatCostUsagebyService := FormatCostUsagebyService(costUsageByService)
	reports.GenerateReport(formatCostUsagebyService)
}
