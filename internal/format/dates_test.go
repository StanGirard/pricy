package format

import "testing"

var datesIntervals DateIntervals = DateIntervals{
	DateInterval{Start: "2018-01-03", End: "2018-01-04"},
	DateInterval{Start: "2018-01-01", End: "2018-01-02"},
	DateInterval{Start: "2018-01-02", End: "2018-01-03"},
	DateInterval{Start: "2018-01-04", End: "2018-01-05"},
	DateInterval{Start: "2018-01-05", End: "2018-01-06"},
	DateInterval{Start: "2018-01-06", End: "2018-01-07"},
}

func TestHeaders(t *testing.T) {
	t.Run("TestHeaders", func(t *testing.T) {
		headers := datesIntervals.Headers()
		if len(headers) != 7 {
			t.Errorf("Expected 7 headers, got %d", len(headers))
		}
	})
}

func TestGetStartDate(t *testing.T) {
	t.Run("TestGetStartDate", func(t *testing.T) {
		dateInterval := datesIntervals[0]
		if dateInterval.GetStartDate() != "2018-01-03" {
			t.Errorf("Expected 2018-01-03, got %s", dateInterval.GetStartDate())
		}
	})
}

func TestGetEndDate(t *testing.T) {
	t.Run("TestGetEndDate", func(t *testing.T) {
		dateInterval := datesIntervals[0]
		if dateInterval.GetEndDate() != "2018-01-04" {
			t.Errorf("Expected 2018-01-04, got %s", dateInterval.GetEndDate())
		}
	})
}

func TestString(t *testing.T) {
	t.Run("TestString", func(t *testing.T) {
		dateInterval := datesIntervals[0]
		if dateInterval.String() != "2018-01-03 - 2018-01-04" {
			t.Errorf("Expected 2018-01-03 - 2018-01-04, got %s", dateInterval.String())
		}
	})
}

func TestContainsDateInterval(t *testing.T) {
	t.Run("TestContainsDateInterval", func(t *testing.T) {
		dateInterval := DateInterval{Start: "2018-01-01", End: "2018-01-02"}
		if !ContainsDateInterval(datesIntervals, dateInterval) {
			t.Errorf("Expected true, got false")
		}
	})
}

func TestSortDates(t *testing.T) {
	t.Run("TestSortDates", func(t *testing.T) {
		dates := SortDates(datesIntervals)
		if dates[0].Start != "2018-01-01" {
			t.Errorf("Expected 2018-01-01, got %s", dates[0].Start)
		}
		if dates[len(dates)-1].End != "2018-01-07" {
			t.Errorf("Expected 2018-01-07, got %s", dates[len(dates)-1].End)
		}
	})
}
