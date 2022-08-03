package format

type Configuration struct {
	Granularity string `json:"granularity,omitempty"`
	Days        int    `json:"days,omitempty"`
	Interval    string `json:"interval,omitempty"`
}
