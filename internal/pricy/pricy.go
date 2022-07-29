package pricy

import (
	"github.com/stangirard/pricy/internal/aws"
	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/reports"
)

type ServicesArray []format.Service

func Run() error {
	cost := aws.InitAWS()
	reports.InitReport(cost)

	return nil
}
