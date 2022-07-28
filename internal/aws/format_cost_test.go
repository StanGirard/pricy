package aws

import (
	"github.com/stangirard/pricy/internal/format"
)

var ServicePrices []format.ServicePrice = []format.ServicePrice{
	{Service: "Amazon Elastic Compute Cloud", Cost: 1.0, Units: "USD"},
	{Service: "EC2", Cost: 1.0, Units: "USD"},
}
