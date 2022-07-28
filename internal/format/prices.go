package format

import "fmt"

type ServicePrice struct {
	Service      string
	Cost         float64
	Units        string
	DateInterval DateInterval
}

var UnitsToSymbol = map[string]string{
	"USD": "$",
	"GBP": "£",
	"EUR": "€",
	"JPY": "¥",
	"CAD": "$",
	"AUD": "$",
	"NZD": "$",
	"CHF": "Fr",
	"SEK": "kr",
	"NOK": "kr",
	"DKK": "kr",
	"KRW": "₩",
	"RUB": "₽",
	"INR": "₹",
	"BRL": "R$",
	"TRY": "₺",
}

// define print function for ServicePrice
func (s ServicePrice) Print() {
	fmt.Printf("Service: %s, Cost: %f%s\n", s.Service, s.Cost, UnitsToSymbol[s.Units])
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
