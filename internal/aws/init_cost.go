package aws

import (
	"flag"

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
	costUsageByService := getCostUsageByService(costExplorer, dateInterval.Start, dateInterval.End)
	formatCostUsagebyService := formatCostUsagebyService(costUsageByService)
	reportGenerate(formatCostUsagebyService)
}
