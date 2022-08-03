package helpers

import (
	"testing"
	"time"
)

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

func TestDaysInterval(t *testing.T) {
	interval := DaysInterval(1)
	if interval.Start != time.Now().AddDate(0, 0, -2).Format("2006-01-02") || interval.End != time.Now().AddDate(0, 0, -1).Format("2006-01-02") {
		t.Errorf("Invalid interval")
	}
}
