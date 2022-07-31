package pricy

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/stangirard/pricy/internal/aws"
	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/prom"
	"github.com/stangirard/pricy/internal/reports"
)

var (
	prometheus = flag.Bool("prometheus", false, "print prometheus metrics")
)

type Services []format.Service

var promServices *[]format.Service

var m sync.Mutex

func Execute() error {
	flag.Parse()
	if *prometheus {
		promServices := &[]format.Service{}
		GetServices(promServices, &m)
		UpdateServices(&promServices, &m)
		prom.Execute(&promServices, &m)
		return nil
	}
	cost := aws.Execute()
	reports.InitReport(cost)

	return nil
}

func GetServices(services *[]format.Service, m *sync.Mutex) {
	m.Lock()
	value := aws.Execute()
	*services = value
	m.Unlock()
}

func UpdateServices(services **[]format.Service, m *sync.Mutex) {
	var timer = time.Now()
	go func() {
		for {
			if time.Since(timer).Seconds() > 10 {
				timer = time.Now()
				GetServices(*services, m)
				newServices := make([]format.Service, len(**services))
				fmt.Println("Time: " + time.Now().Format("2006-01-02 15:04:05"))
				m.Lock()
				for _, service := range **services {
					dates := format.SortDates(format.FindDatesIntervals(**services))
					newService := service
					newService.DatePrice[dates[0]] = 0.0
					newServices = append(newServices, newService)
					fmt.Println(newService.Name + ": " + fmt.Sprintf("%f", service.DatePrice[dates[0]]))
				}
				*services = &newServices
				m.Unlock()

			}
			time.Sleep(1 * time.Second)
			fmt.Println("Time since last update: " + time.Since(timer).String())
		}
	}()

}
