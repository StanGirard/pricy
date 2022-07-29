package format

type DateInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type DateIntervals []DateInterval

type Service struct {
	Service        string
	Units          string
	Account        string
	DatePrice      map[DateInterval]float64
	PriceEvolution map[DateInterval]float64
}

type Services []Service

type TotalPerDay struct {
	Date      DateInterval
	TotalCost float64
}
