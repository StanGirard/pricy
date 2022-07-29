package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/stangirard/pricy/internal/format"
)

func DaysInterval(days int) format.DateInterval {
	start := time.Now().AddDate(0, 0, -days-1)
	end := time.Now().AddDate(0, 0, -1)
	return format.DateInterval{
		Start: start.Format("2006-01-02"),
		End:   end.Format("2006-01-02"),
	}
}

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
