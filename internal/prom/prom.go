package prom

import (
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stangirard/pricy/internal/format"
)

func cleanString(str string) string {
	// replace every non-alphanumeric character with an underscore
	str = strings.Replace(str, " ", "_", -1)
	str = strings.Replace(str, ".", "", -1)
	str = strings.Replace(str, ":", "", -1)
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, ";", "", -1)
	str = strings.Replace(str, "!", "", -1)
	str = strings.Replace(str, "?", "", -1)
	str = strings.Replace(str, "\"", "", -1)
	str = strings.Replace(str, "'", "", -1)
	str = strings.Replace(str, "`", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	str = strings.Replace(str, "{", "", -1)
	str = strings.Replace(str, "}", "", -1)
	str = strings.Replace(str, "=", "", -1)
	str = strings.Replace(str, "-", "_", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.ToLower(str)
	space := regexp.MustCompile(`_+`)
	str = space.ReplaceAllString(str, "_")

	return str

}

var gauges = make(map[string]prometheus.Gauge)

func recordMetrics(services *[]format.Service, m *sync.Mutex) {

	flag.Set("days", "1")
	var timer = time.Now()
	m.Lock()
	for _, service := range *services {
		//Get Dates
		// Get first date in Service.DatePrice map
		var firstDate format.DateInterval
		for date := range service.DatePrice {
			firstDate = date
			break
		}
		gauges[service.Name] = promauto.NewGauge(prometheus.GaugeOpts{
			Name: "pricy_" + cleanString(service.Name),
			Help: "Price of " + service.Name,
		})
		gauges[service.Name].Set(service.DatePrice[firstDate])
	}
	m.Unlock()
	go func() {
		for {

			if time.Since(timer).Seconds() > 60 {
				m.Lock()
				fmt.Println("Updating Metrics at " + time.Now().Format("2006-01-02 15:04:05"))
				timer = time.Now()
				for _, service := range *services {
					// Update prometheus gauge
					var firstDate format.DateInterval
					for date := range service.DatePrice {
						firstDate = date
						break
					}

					gauges[service.Name].Set(service.DatePrice[firstDate])

				}
				m.Unlock()
				time.Sleep(10 * time.Second)

			}
		}
	}()
}

func Execute(services *[]format.Service, m *sync.Mutex) {
	recordMetrics(services, m)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
