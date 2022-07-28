package format

import "fmt"

type ServicePrice struct {
	Service string
	Cost    float64
	Units   string
}

// define print function for ServicePrice
func (s ServicePrice) Print() {
	fmt.Printf("Service: %s, Cost: %f, Units: %s\n", s.Service, s.Cost, s.Units)
}
