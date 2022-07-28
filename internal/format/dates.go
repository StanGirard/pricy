package format

import "fmt"

type DateInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (d DateInterval) GetStartDate() string {
	return d.Start
}

func (d DateInterval) GetEndDate() string {
	return d.End
}

func (d DateInterval) Print() {
	fmt.Printf("Start: %s, End: %s\n", d.Start, d.End)
}
