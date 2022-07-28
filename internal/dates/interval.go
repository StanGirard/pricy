package dates

import (
	"time"

	"github.com/stangirard/pricy/internal/format"
)

// Functions that returns a date interval from date and week before
func WeekInterval(date time.Time) format.DateInterval {
	start := date.AddDate(0, 0, -7)
	end := date
	return format.DateInterval{
		Start: start.Format("2006-01-02"),
		End:   end.Format("2006-01-02"),
	}

}

func FromLastWeekToNow() format.DateInterval {
	return WeekInterval(time.Now())
}

func FromLastMonthToNow() format.DateInterval {
	start := time.Now().AddDate(0, -1, 0)
	end := time.Now()
	return format.DateInterval{
		Start: start.Format("2006-01-02"),
		End:   end.Format("2006-01-02"),
	}
}
