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
	date        = flag.String("date", "", "date to get cost usage")
	month       = flag.Bool("month", false, "get cost usage for month")
	granularity = flag.String("granularity", "DAILY", "granularity to get cost usage")
)

func createCostExplorer(sess *session.Session) *costexplorer.CostExplorer {
	return costexplorer.New(sess)
}

func InitCostExplorer(session *session.Session) {
	// Initialize the session
	costExplorer := createCostExplorer(session)

	// Generating Date
	var dateInterval format.DateInterval
	if *month {
		dateInterval = helpers.FromLastMonthToNow()
	} else {
		dateInterval = helpers.FromLastWeekToNow()
	}
	// uppercase string for the granularity
	costUsageByService := getCostUsageByService(costExplorer, dateInterval.Start, dateInterval.End, strings.ToUpper(*granularity))
	formatCostUsagebyService := FormatCostUsagebyService(costUsageByService)
	reports.GenerateReport(formatCostUsagebyService)
}
