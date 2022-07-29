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

func GenerateCostCSVArray(Services []format.Service) [][]string {
	var csvArray [][]string
	csvArray = append(csvArray, format.FindDatesIntervals(Services).Headers())
	dates := format.SortDates(format.FindDatesIntervals(Services))
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

func WriteCSV(csvArray [][]string, filepath string) {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	writeError := w.WriteAll(csvArray)
	if writeError != nil {
		log.Fatalln("failed to write to csv file", err)
	}

}

func CSVEvolutionReport(Services []format.Service) [][]string {
	var csvArray [][]string
	csvArray = append(csvArray, format.FindDatesIntervals(Services).Headers())
	dates := format.SortDates(format.FindDatesIntervals(Services))
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

func (services ServicesArray) initCSV() {
	flag.Parse()

	if *csvFlag {
		csvArray := GenerateCostCSVArray(services)
		WriteCSV(csvArray, "reports.csv")
		if *evolution {
			evolutionCSVArray := CSVEvolutionReport(services)
			WriteCSV(evolutionCSVArray, "evolution.csv")
		}
	}
}
