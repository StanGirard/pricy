package format

import "testing"

var datesIntervals DateIntervals = DateIntervals{
	DateInterval{Start: "2018-01-01", End: "2018-01-02"},
	DateInterval{Start: "2018-01-02", End: "2018-01-03"},
	DateInterval{Start: "2018-01-03", End: "2018-01-04"},
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
