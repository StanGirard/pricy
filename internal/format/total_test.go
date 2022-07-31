package format

import "testing"

func TestTotalCostUsage(t *testing.T) {
	var service Service
	service.Name = "Amazon Elastic Compute Cloud"
	service.Units = "USD"
	service.DatePrice = map[DateInterval]float64{
		{Start: "2019-01-01", End: "2019-01-31"}: 0.1,
		{Start: "2019-02-01", End: "2019-02-28"}: 0.1,
		{Start: "2019-03-01", End: "2019-03-31"}: 0.1,
	}
	service.PriceEvolution = map[DateInterval]float64{
		{Start: "2019-01-01", End: "2019-01-31"}: 0.0,
		{Start: "2019-02-01", End: "2019-02-28"}: 0.02,
		{Start: "2019-03-01", End: "2019-03-31"}: 0.03,
	}

	var service2 Service
	service2.Name = "Amazon Elastic Compute Cloud 2"
	service2.Units = "USD"
	service2.DatePrice = map[DateInterval]float64{
		{Start: "2019-01-01", End: "2019-01-31"}: 0.1,
		{Start: "2019-02-01", End: "2019-02-28"}: 0.1,
		{Start: "2019-03-01", End: "2019-03-31"}: 0.1,
	}
	service2.PriceEvolution = map[DateInterval]float64{
		{Start: "2019-01-01", End: "2019-01-31"}: 0.0,
		{Start: "2019-02-01", End: "2019-02-28"}: 0.02,
		{Start: "2019-03-01", End: "2019-03-31"}: 0.03,
	}

	services := []Service{service, service2}
	totalCost := TotalCostUsage(services)
	for _, totalPerDay := range totalCost {
		if totalPerDay.TotalCost != 0.2 {
			t.Errorf("Expected 0.1, got %f", totalPerDay.TotalCost)
		}
	}
}
