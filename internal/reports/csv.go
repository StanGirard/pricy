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
	csvFlag   = flag.Bool("csv", false, "print in csv format")
	evolution = flag.Bool("evolution", false, "print evolution report")
)

func GenerateHeader(dates []format.DateInterval) []string {
	var header []string
	header = append(header, "Service")
	var datesHeader []string
	for _, date := range dates {
		datesHeader = append(datesHeader, date.Start)
	}
	sort.Slice(datesHeader, func(i, j int) bool {
		return datesHeader[i] < datesHeader[j]
	})
	header = append(header, datesHeader...)
	return header
}

func GenerateCostCSVArray(Services []format.Service) [][]string {
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

func WriteCSVToFile(csvArray [][]string, filepath string) {
	f, err := os.Create(filepath)
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(csvArray)

}

func GenerateCSVReport(Services []format.Service) {
	flag.Parse()

	if *csvFlag {
		csvArray := GenerateCostCSVArray(Services)
		WriteCSVToFile(csvArray, "reports.csv")
		if *evolution {
			evolutionCSVArray := GenerateCSVEvolutionReport(Services)
			WriteCSVToFile(evolutionCSVArray, "evolution.csv")
		}
	}
}

func GenerateCSVEvolutionReport(Services []format.Service) [][]string {
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
			csvRow = append(csvRow, fmt.Sprintf("%.2f", service.PriceEvolution[date]))
		}
		csvArray = append(csvArray, csvRow)
	}
	return csvArray
}
