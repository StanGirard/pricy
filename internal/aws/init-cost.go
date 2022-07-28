package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stangirard/pricy/internal/dates"
	"github.com/stangirard/pricy/internal/format"
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
	costUsageByServiceOutput := getCostUsageByService(costExplorer, dateInterval.Start, dateInterval.End)
	fmt.Println("Total Cost Usage: ", TotalCostUsage(costUsageByServiceOutput), "From:", dateInterval.Start, "To:", dateInterval.End)
}
