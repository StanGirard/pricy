package reports

import (
	"fmt"
	"testing"

	"github.com/stangirard/pricy/internal/format"
)

func TestCalculateEvolution(t *testing.T) {
	services := Services{}
	service.PriceEvolution = map[format.DateInterval]float64{}
	service2.PriceEvolution = map[format.DateInterval]float64{}
	services = append(services, service)
	services = append(services, service2)
	services.calculateEvolution()
	fmt.Println(services)
	for _, service := range services {
		for i, value := range service.PriceEvolution {
			if i.Start == "2019-01-01" && i.End == "2019-01-31" && value == 0.0 {
				continue
			} else if i.Start == "2019-02-01" && i.End == "2019-02-28" && value == 0.00 {
				continue
			} else if i.Start == "2019-03-01" && i.End == "2019-03-31" && value == 0.00 {
				continue
			} else {
				t.Errorf("Expected 0.00, got %f", value)
			}
		}
	}
}
