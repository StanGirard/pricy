package format

// Struct that contains a start and an end string usually formated as follows : "YYYY-MM-DD"
type DateInterval struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

// Array of DateInterval
type DateIntervals []DateInterval

// The backbone interface of the projet. All cloud providers output this object and is used by all other packages
type Service struct {
	Name           string                   `json:"name,omitempty"`
	Units          string                   `json:"units,omitempty"`
	Account        string                   `json:"account,omitempty"`
	DatePrice      map[DateInterval]float64 `json:"date_price,omitempty"`
	PriceEvolution map[DateInterval]float64 `json:"price_evolution,omitempty"`
}

// Array of services
type Services []Service

// Struct to store the total per date interval
type TotalPerDay struct {
	Date      DateInterval `json:"date,omitempty"`
	TotalCost float64      `json:"total_cost,omitempty"`
}
