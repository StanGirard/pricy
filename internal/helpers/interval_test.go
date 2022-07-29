package helpers

import "testing"

func TestParseInterval(t *testing.T) {
	interval, err := ParseInterval("2020-01-01:2020-01-02")
	if interval.Start != "2020-01-01" || interval.End != "2020-01-02" || err != nil {
		t.Errorf("Invalid interval")
	}
	interval, err = ParseInterval("2020-01-01")
	if err == nil {
		t.Errorf("Invalid interval")
	}
}
