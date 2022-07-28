package format

import "fmt"

type ServicePrice struct {
	Service string
	Cost    float64
	Units   string
}

type PricePerDate struct {
	DateInterval DateInterval
	ServicePrice []ServicePrice
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
