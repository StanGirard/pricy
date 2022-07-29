package format

type Service struct {
	Service        string
	Units          string
	Account        string
	DatePrice      map[DateInterval]float64
	PriceEvolution map[DateInterval]float64
}

func AddServices(servicesArray []Service, service, units string, Date DateInterval, price float64) []Service {
	for _, serv := range servicesArray {
		if serv.Service == service {
			serv.DatePrice[Date] = price
			return servicesArray
		}
	}
	servicesArray = append(servicesArray, Service{
		Service: service,
		Units:   units,
		DatePrice: map[DateInterval]float64{
			Date: price,
		},
		PriceEvolution: map[DateInterval]float64{},
	})
	return servicesArray
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

func FindAllDateIntervals(services []Service) []DateInterval {
	var dates []DateInterval
	for _, service := range services {
		for date := range service.DatePrice {
			if !ContainsDateInterval(dates, date) {
				dates = append(dates, date)
			}
		}
	}
	return dates
}
