package format

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
