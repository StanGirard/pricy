package reports

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/stangirard/pricy/internal/format"
)

var (
	csvFlag = flag.Bool("csv", false, "print in csv format")
)

func GenerateHeader(dates []format.DateInterval) []string {
	var header []string
	header = append(header, "Service")
	for _, date := range dates {
		header = append(header, date.Start)
	}
	return header
}

func GenerateCSVArray(Services []format.Service) [][]string {
	var csvArray [][]string
	csvArray = append(csvArray, GenerateHeader(format.FindAllDateIntervals(Services)))
	dates := format.SortDates(format.FindAllDateIntervals(Services))
	sort.Slice(Services, func(i, j int) bool {
		return Services[i].Service < Services[j].Service
	})
	for _, service := range Services {
		var csvRow []string
		csvRow = append(csvRow, service.Service)
		for _, date := range dates {
			csvRow = append(csvRow, fmt.Sprintf("%.2f", service.DatePrice[date]))
		}
		csvArray = append(csvArray, csvRow)
	}
	return csvArray
}

func WriteCSVToFile(csvArray [][]string) {
	f, err := os.Create("reports.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(csvArray)

}

func GenerateCSVReport(Services []format.Service) {
	flag.Parse()
	csvArray := GenerateCSVArray(Services)
	if *csvFlag {
		WriteCSVToFile(csvArray)
	}
}
