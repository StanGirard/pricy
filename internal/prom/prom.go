package prom

import (
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stangirard/pricy/internal/aws"
	"github.com/stangirard/pricy/internal/format"
)

var (
// map of gauge metrics with service name as key

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
	str = strings.ToLower(str)
	space := regexp.MustCompile(`_+`)
	str = space.ReplaceAllString(str, "_")

	return str

}

func recordMetrics() {
	services := aws.InitAWS()
	gauges := make(map[string]prometheus.Gauge)
	flag.Set("days", "1")
	var timer = time.Now()

	for _, service := range services {
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

	go func() {
		for {
			if time.Since(timer).Seconds() > 28800 {
				timer = time.Now()
				services = aws.InitAWS()
				fmt.Println("Time: " + time.Now().Format("2006-01-02 15:04:05"))
				for _, service := range services {
					// Update prometheus gauge
					var firstDate format.DateInterval
					for date := range service.DatePrice {
						firstDate = date
						break
					}
					gauges[service.Name].Set(service.DatePrice[firstDate])

					fmt.Println("Updated gauge for " + cleanString(service.Name))

				}

			}
			time.Sleep(10 * time.Second)
			fmt.Println("Time since last update: " + time.Since(timer).String())
		}
	}()
}

func RunProm() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
