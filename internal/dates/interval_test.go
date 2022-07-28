package dates

import (
	"testing"
	"time"

	"github.com/stangirard/pricy/internal/format"
	"github.com/stretchr/testify/assert"
)

// create a date that is 2022-01-01
var date = time.Date(2022, 2, 1, 0, 0, 1, 0, time.UTC)

func TestWeekInterval(t *testing.T) {
	assert.Equal(t, format.DateInterval{Start: "2022-01-25", End: "2022-02-01"}, WeekInterval(date))
}

func TestMonthInterval(t *testing.T) {
	assert.Equal(t, format.DateInterval{Start: "2022-01-01", End: "2022-02-01"}, MonthInterval(date))
}
