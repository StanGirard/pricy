package cmd

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

var m sync.Mutex

func Execute() error {
	flag.Parse()
	flag.Set("days", "1")
	services := make([]format.Service, 0)
	getServices(&services, &m)

	if *prometheus {
		updateServices(&services, &m)
		prom.Execute(&services, &m)
		return nil
	}
	reports.InitReport(services)

	return nil
}

func getServices(services *[]format.Service, m *sync.Mutex) {
	m.Lock()
	value := aws.Execute()
	*services = value
	m.Unlock()
}

func updateServices(services *[]format.Service, m *sync.Mutex) {
	var timer = time.Now()
	go func() {
		for {
			if time.Since(timer).Seconds() > 28800 {
				timer = time.Now()
				m.Lock()
				getServices(services, m)
				m.Unlock()
				fmt.Println("Time Update: " + time.Now().Format("2006-01-02 15:04:05"))
			}
			time.Sleep(1 * time.Second)
		}
	}()

}
