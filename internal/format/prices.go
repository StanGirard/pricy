package format

import "fmt"

type ServicePrice struct {
	Service      string
	Cost         float64
	Units        string
	DateInterval DateInterval
}

// define print function for ServicePrice
func (s ServicePrice) Print() {
	fmt.Printf("Service: %s, Cost: %f, Units: %s\n", s.Service, s.Cost, s.Units)
}

func (s ServicePrice) PrintDateInterval() {
	fmt.Printf("Start: %s, End: %s\n", s.DateInterval.Start, s.DateInterval.End)
}

func (s ServicePrice) getStartDate() string {
	return s.DateInterval.Start
}

func (s ServicePrice) getEndDate() string {
	return s.DateInterval.End
}
