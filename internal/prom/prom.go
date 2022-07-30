package prom

import (
	"flag"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stangirard/pricy/internal/aws"
)

var (
	opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "my_gauge",
		Help: "A simple example of a Gauge metric",
	})
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
	flag.Set("days", "1")

	for _, service := range services {
		//Get Dates
		promauto.NewGauge(prometheus.GaugeOpts{
			Name: "pricy_" + cleanString(service.Name),
			Help: "Price of " + service.Name,
		})
	}

	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func RunProm() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
