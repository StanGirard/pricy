package reports

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/stangirard/pricy/internal/format"
	"github.com/stangirard/pricy/internal/gsheet"
)

var (
	csvFlag    = flag.Bool("csv", false, "print in csv format")
	evolution  = flag.Bool("evolution", false, "print evolution report")
	gsheetFlag = flag.Bool("gsheet", false, "print google sheet metrics")
)

// Takes an array of services and turns it into an array of array.
func generateCostCSVArray(Services []format.Service) [][]string {
	var csvArray [][]string
	csvArray = append(csvArray, format.FindDatesIntervals(Services).Headers())
	dates := format.SortDates(format.FindDatesIntervals(Services))
	sort.Slice(Services, func(i, j int) bool {
		return Services[i].Name < Services[j].Name
	})
	for _, service := range Services {
		var csvRow []string
		csvRow = append(csvRow, service.Name)
		for _, date := range dates {
			csvRow = append(csvRow, fmt.Sprintf("%.2f", service.DatePrice[date]))
		}
		csvArray = append(csvArray, csvRow)
	}
	return csvArray
}

// Write an array of arrays into a csv file
func writeCSV(csvArray [][]string, filepath string) {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	writeError := w.WriteAll(csvArray)
	if writeError != nil {
		log.Fatalln("failed to write to csv file", err)
	}

}

// Generate an array of arrays for the evolutions
func csvEvolutionReport(Services []format.Service) [][]string {
	var csvArray [][]string
	csvArray = append(csvArray, format.FindDatesIntervals(Services).Headers())
	dates := format.SortDates(format.FindDatesIntervals(Services))
	sort.Slice(Services, func(i, j int) bool {
		return Services[i].Name < Services[j].Name
	})
	for _, service := range Services {
		var csvRow []string
		csvRow = append(csvRow, service.Name)
		for _, date := range dates {
			csvRow = append(csvRow, fmt.Sprintf("%.2f", service.PriceEvolution[date]))
		}
		csvArray = append(csvArray, csvRow)
	}
	return csvArray
}

// Initis the csv service with the logic around the flags
func (services Services) initCSV() {
	flag.Parse()

	if *csvFlag {
		csvArray := generateCostCSVArray(services)
		writeCSV(csvArray, "reports.csv")
		if *evolution {
			evolutionCSVArray := csvEvolutionReport(services)
			writeCSV(evolutionCSVArray, "evolution.csv")
		}
		if *gsheetFlag {
			gsheet.Execute(csvArray)
		}
	}
}
