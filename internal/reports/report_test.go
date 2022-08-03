package reports

import "testing"

func TestGetCostByDate(t *testing.T) {
	var pricetoDates priceToDateArray
	pricetoDates = append(pricetoDates, priceToDate{Name: "service1", Price: 1.0})
	pricetoDates = append(pricetoDates, priceToDate{Name: "service2", Price: 2.0})
	pricetoDates = append(pricetoDates, priceToDate{Name: "service3", Price: 3.0})
	pricetoDates = append(pricetoDates, priceToDate{Name: "service4", Price: 4.0})

	t.Run("TestGetCostByDate", func(t *testing.T) {
		value := pricetoDates.getCostByDate()
		if value != 10.0 {
			t.Errorf("Expected 10.0, got %f", value)
		}
	})

}
