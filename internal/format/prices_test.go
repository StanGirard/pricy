package format

import (
	"sort"
	"testing"
)

var mockServices Services = []Service{
	{
		Service: "Amazon",
		Units:   "USD",
		DatePrice: map[DateInterval]float64{
			{Start: "2019-01-01", End: "2019-01-31"}: 10.0,
			{Start: "2019-02-01", End: "2019-02-28"}: 20.0,
			{Start: "2019-03-01", End: "2019-03-31"}: 30.0,
			{Start: "2019-04-01", End: "2019-04-30"}: 40.0,
		},
		PriceEvolution: map[DateInterval]float64{},
	}, {
		Service: "Amazon2",
		Units:   "USD",
		DatePrice: map[DateInterval]float64{
			{Start: "2019-01-01", End: "2019-01-31"}: 10.0,
			{Start: "2019-02-01", End: "2019-02-28"}: 20.0,
			{Start: "2019-03-01", End: "2019-03-31"}: 30.0,
			{Start: "2019-04-01", End: "2019-04-30"}: 40.0,
		},
		PriceEvolution: map[DateInterval]float64{},
	},
}

func TestAddServices(t *testing.T) {
	testServices := mockServices
	testServices = AddServices(testServices, "Amazon3", "USD", DateInterval{Start: "2019-01-01", End: "2019-01-31"}, 10.0)
	if len(testServices) != 3 {
		t.Errorf("Expected 3 services, got %d", len(testServices))
	}
	testServices = AddServices(testServices, "Amazon3", "USD", DateInterval{Start: "2019-02-01", End: "2019-02-28"}, 10.0)
	if len(testServices) != 3 {
		t.Errorf("Expected 3 services, got %d", len(testServices))
	}
}

func TestFindDatesIntervals(t *testing.T) {
	testServices := mockServices
	dates := FindDatesIntervals(testServices)
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Start < dates[j].Start
	})
	if len(dates) != 4 {
		t.Errorf("Expected 4 dates, got %d", len(dates))
	}
	if dates[0].Start != "2019-01-01" || dates[0].End != "2019-01-31" {
		t.Errorf("Expected dates[0] to be 2019-01-01 to 2019-01-31, got %s to %s", dates[0].Start, dates[0].End)
	}
	if dates[1].Start != "2019-02-01" || dates[1].End != "2019-02-28" {
		t.Errorf("Expected dates[1] to be 2019-02-01 to 2019-02-28, got %s to %s", dates[1].Start, dates[1].End)
	}
	if dates[2].Start != "2019-03-01" || dates[2].End != "2019-03-31" {
		t.Errorf("Expected dates[2] to be 2019-03-01 to 2019-03-31, got %s to %s", dates[2].Start, dates[2].End)
	}
	if dates[3].Start != "2019-04-01" || dates[3].End != "2019-04-30" {
		t.Errorf("Expected dates[3] to be 2019-04-01 to 2019-04-30, got %s to %s", dates[3].Start, dates[3].End)
	}

}
