package format

func TotalCostUsage(Services Services) []TotalPerDay {
	var TotalPerDays []TotalPerDay
	for _, service := range Services {
		for i, value := range service.DatePrice {
			var totalPerDay TotalPerDay
			totalPerDay.Date = DateInterval{Start: i.Start, End: i.End}
			totalPerDay.TotalCost = value
			already := false
			for j, value := range TotalPerDays {
				if totalPerDay.Date == value.Date {
					TotalPerDays[j].TotalCost += totalPerDay.TotalCost
					already = true
					break
				}
			}
			if !already {
				TotalPerDays = append(TotalPerDays, totalPerDay)
			}
		}
	}
	return TotalPerDays
}
