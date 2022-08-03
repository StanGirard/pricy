package format

// Configuration that can be passed to many packages in order to pass the flags down without cyclic dependencies
type Configuration struct {
	Granularity string `json:"granularity,omitempty"`
	Days        int    `json:"days,omitempty"`
	Interval    string `json:"interval,omitempty"`
}
