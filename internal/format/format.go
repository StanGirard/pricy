package format

type DateInterval struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type DateIntervals []DateInterval

type Service struct {
	Name           string                   `json:"name,omitempty"`
	Units          string                   `json:"units,omitempty"`
	Account        string                   `json:"account,omitempty"`
	DatePrice      map[DateInterval]float64 `json:"date_price,omitempty"`
	PriceEvolution map[DateInterval]float64 `json:"price_evolution,omitempty"`
}

type Services []Service

type TotalPerDay struct {
	Date      DateInterval `json:"date,omitempty"`
	TotalCost float64      `json:"total_cost,omitempty"`
}
