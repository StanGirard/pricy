package aws

import (
	"testing"

	"github.com/stangirard/pricy/internal/format"
	"github.com/stretchr/testify/assert"
)

var ServicePrices []format.ServicePrice = []format.ServicePrice{
	{Service: "Amazon Elastic Compute Cloud", Cost: 1.0, Units: "USD"},
	{Service: "EC2", Cost: 1.0, Units: "USD"},
}

func TestTotalCostUsage(t *testing.T) {
	assert.Equal(t, 2.0, TotalCostUsage(ServicePrices))
}
