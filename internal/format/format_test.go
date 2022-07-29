package format

var ServicePrices []Service = []Service{
	{
		Service: "Amazon Elastic Compute Cloud",
		Units:   "USD",
		DatePrice: map[DateInterval]float64{
			{Start: "2019-01-01", End: "2019-01-31"}: 0.0,
			{Start: "2019-02-01", End: "2019-02-28"}: 0.0,
			{Start: "2019-03-01", End: "2019-03-31"}: 0.0,
		},
	},
}
