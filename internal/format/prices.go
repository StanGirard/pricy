package format

// Cloud providers usually return the units as USD. We want to pretty print those into $. Mapping table.
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

// Adds a service in an array with some logic. For example if the name already exist it adds it to the specific service.
// otherwise it creates a new object
func AddServices(Services []Service, service, units string, Date DateInterval, price float64) []Service {
	for _, serv := range Services {
		if serv.Name == service {
			serv.DatePrice[Date] = price
			return Services
		}
	}
	Services = append(Services, Service{
		Name:  service,
		Units: units,
		DatePrice: map[DateInterval]float64{
			Date: price,
		},
		PriceEvolution: map[DateInterval]float64{},
	})
	return Services
}

// Find all the dates intervals in a service array
func FindDatesIntervals(services []Service) DateIntervals {
	var dates DateIntervals
	for _, service := range services {
		for date := range service.DatePrice {
			if !ContainsDateInterval(dates, date) {
				dates = append(dates, date)
			}
		}
	}
	return dates
}
