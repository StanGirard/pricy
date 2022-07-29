package format

func TotalCostUsage(Services []Service) []TotalPerDay {
	var TotalPerDays []TotalPerDay
	for _, service := range Services {
		for date, price := range service.DatePrice {
			for i := range TotalPerDays {
				if TotalPerDays[i].Date.Start == date.Start && TotalPerDays[i].Date.End == date.End {
					TotalPerDays[i].TotalCost += price
					break
				} else {
					TotalPerDays = append(TotalPerDays, TotalPerDay{Date: date, TotalCost: price})
					break
				}
			}
		}
	}
	return TotalPerDays
}
