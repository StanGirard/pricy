package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	t.Run("TestConfiguration", func(t *testing.T) {
		config := Configuration{
			Granularity: "day",
			Days:        30,
			Interval:    "1h",
		}
		assert.Equal(t, config.Granularity, "day")
		assert.Equal(t, config.Days, 30)
		assert.Equal(t, config.Interval, "1h")
	})

}
