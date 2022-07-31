package format

import "testing"

var ServicePrices Services = []Service{
	{
		Name:  "Amazon Elastic Compute Cloud",
		Units: "USD",
		DatePrice: map[DateInterval]float64{
			{Start: "2019-01-01", End: "2019-01-31"}: 0.0,
			{Start: "2019-02-01", End: "2019-02-28"}: 0.0,
			{Start: "2019-03-01", End: "2019-03-31"}: 0.0,
		},
		PriceEvolution: map[DateInterval]float64{
			{Start: "2019-01-01", End: "2019-01-31"}: 0.0,
			{Start: "2019-02-01", End: "2019-02-28"}: 0.02,
			{Start: "2019-03-01", End: "2019-03-31"}: 0.03,
		},
	},
}

func TestServicesPrices(t *testing.T) {
	if len(ServicePrices) != 1 {
		t.Errorf("Expected 1 service, got %d", len(ServicePrices))
	}
	for _, service := range ServicePrices {
		if len(service.DatePrice) != 3 {
			t.Errorf("Expected 3 dates, got %d", len(service.DatePrice))
		}
		if len(service.PriceEvolution) != 3 {
			t.Errorf("Expected 3 dates, got %d", len(service.PriceEvolution))
		}
		if service.DatePrice[DateInterval{Start: "2019-01-01", End: "2019-01-31"}] != 0.0 {
			t.Errorf("Expected 0.0, got %f", service.DatePrice[DateInterval{Start: "2019-01-01", End: "2019-01-31"}])
		}
		if service.PriceEvolution[DateInterval{Start: "2019-01-01", End: "2019-01-31"}] != 0.0 {
			t.Errorf("Expected 0.0, got %f", service.PriceEvolution[DateInterval{Start: "2019-01-01", End: "2019-01-31"}])
		}
	}
}

func TestTypeService(t *testing.T) {
	var service Service
	if service.Name != "" {
		t.Errorf("Expected empty string, got %s", service.Name)
	}
	if service.Units != "" {
		t.Errorf("Expected empty string, got %s", service.Units)
	}
	if len(service.DatePrice) != 0 {
		t.Errorf("Expected empty map, got %d", len(service.DatePrice))
	}
	if len(service.PriceEvolution) != 0 {
		t.Errorf("Expected empty map, got %d", len(service.PriceEvolution))
	}
}
