package cmd

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/stangirard/pricy/internal/aws"
	"github.com/stangirard/pricy/internal/azure"
	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/prom"
	"github.com/stangirard/pricy/internal/reports"
)

var (
	prometheus  = flag.Bool("prometheus", false, "print prometheus metrics")                                        // Creates a prometheus exporter on http://localhost:2112/metrics
	azureFlag   = flag.Bool("azure", false, "print azure metrics")                                                  // Get the cost from azure
	awsFlag     = flag.Bool("aws", false, "print aws metrics")                                                      // Get the cost from aws
	granularity = flag.String("granularity", "DAILY", "granularity to get cost usage")                              // Granularity to get cost usage, either DAILY or MONTHLY
	days        = flag.Int("days", 14, "get cost usage for last 14 days")                                           // Time range in days since today
	interval    = flag.String("interval", "", "get cost usage for a specific interval as '2022-03-30:2022-03-31' ") // Date interval for the cost usage, if empty, it will get the cost usage for the last 14 days
)

type Services []format.Service

var m sync.Mutex // Mutex to protect write to the services variable when using prometheus

// Executes Pricy
func Execute() error {
	flag.Parse()
	configuration := format.Configuration{
		Granularity: *granularity,
		Days:        *days,
		Interval:    *interval,
	}

	services := make([]format.Service, 0)
	getServices(&services, &m, configuration)

	if *prometheus {
		flag.Set("days", "1")
		updateServices(&services, &m, configuration)
		prom.Execute(&services, &m)
		return nil
	}
	reports.InitReport(services)

	return nil
}

// Get the cost of all services from the cloud provider that is specified with the flag.
func getServices(services *[]format.Service, m *sync.Mutex, configuration format.Configuration) {
	m.Lock()
	if *azureFlag {
		value := azure.Execute(configuration)
		*services = value
	} else if *awsFlag {
		value := aws.Execute(configuration)
		*services = value
	} else {
		panic("No provider selected (--azure or --aws)")
	}
	m.Unlock()
}

// Updates the services price every 8 hours
func updateServices(services *[]format.Service, m *sync.Mutex, configuration format.Configuration) {
	var timer = time.Now()
	go func() {
		for {
			if time.Since(timer).Seconds() > 28800 {
				timer = time.Now()
				m.Lock()
				getServices(services, m, configuration)
				m.Unlock()
				fmt.Println("Time Update: " + time.Now().Format("2006-01-02 15:04:05"))
			}
			time.Sleep(1 * time.Second)
		}
	}()

}
