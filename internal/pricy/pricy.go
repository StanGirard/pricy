package pricy

import (
	"flag"

	"github.com/stangirard/pricy/internal/aws"
	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/prom"
	"github.com/stangirard/pricy/internal/reports"
)

var (
	prometheus = flag.Bool("prometheus", false, "print prometheus metrics")
)

type Services []format.Service

func Run() error {
	flag.Parse()
	if *prometheus {
		prom.RunProm()
		return nil
	}
	cost := aws.InitAWS()
	reports.InitReport(cost)

	return nil
}
