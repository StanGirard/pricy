package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/stangirard/pricy/internal/format"
)

// Gives you a date interval from a number of days before today
// Start at yesterday because the cost usage is for the previous day
func DaysInterval(days int) format.DateInterval {
	start := time.Now().AddDate(0, 0, -days-1)
	end := time.Now().AddDate(0, 0, -1)
	return format.DateInterval{
		Start: start.Format("2006-01-02"),
		End:   end.Format("2006-01-02"),
	}
}

// Parse the date interval from a string like this "2022-03-30:2022-03-31"
func ParseInterval(interval string) (format.DateInterval, error) {
	split := strings.Split(interval, ":")
	if len(split) != 2 {
		return format.DateInterval{}, fmt.Errorf("Invalid interval")
	}
	return format.DateInterval{
		Start: split[0],
		End:   split[1],
	}, nil
}
