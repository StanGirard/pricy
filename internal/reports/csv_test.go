package reports

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/stangirard/pricy/internal/format"
)

var service = format.Service{
	Name:    "Amazon Elastic Compute Cloud",
	Units:   "USD",
	Account: "123456789",
	DatePrice: map[format.DateInterval]float64{
		{Start: "2019-01-01", End: "2019-01-31"}: 0.1,
		{Start: "2019-02-01", End: "2019-02-28"}: 0.0,
		{Start: "2019-03-01", End: "2019-03-31"}: 0.0,
		{Start: "2019-04-01", End: "2019-04-30"}: 0.1,
		{Start: "2019-05-01", End: "2019-05-31"}: 0.2,
	},
	PriceEvolution: map[format.DateInterval]float64{},
}
var service2 = format.Service{
	Name:    "Amazon Elastic Compute Cloud 2",
	Units:   "USD",
	Account: "123456789",
	DatePrice: map[format.DateInterval]float64{
		{Start: "2019-01-01", End: "2019-01-31"}: 0.1,
		{Start: "2019-02-01", End: "2019-02-28"}: 0.0,
		{Start: "2019-03-01", End: "2019-03-31"}: 0.0,
		{Start: "2019-04-01", End: "2019-04-30"}: 0.1,
		{Start: "2019-05-01", End: "2019-05-31"}: 0.2,
	},
	PriceEvolution: map[format.DateInterval]float64{},
}

func TestGenerateCostCSVArray(t *testing.T) {
	var services format.Services
	services = append(services, service)
	services = append(services, service2)
	csvArray := generateCostCSVArray(services)
	t.Run("TestSize", func(t *testing.T) {
		if len(csvArray) != 3 {
			t.Errorf("Expected 3, got %d", len(csvArray))
		}
	})

}

func TestCsvEvolutionReport(t *testing.T) {
	var services format.Services
	services = append(services, service)
	services = append(services, service2)
	csvArray := csvEvolutionReport(services)
	t.Run("TestSize", func(t *testing.T) {
		if len(csvArray) != 3 {
			t.Errorf("Expected 3, got %d", len(csvArray))
		}
	})
}

func TestWriteCSV(t *testing.T) {
	var services format.Services
	services = append(services, service)
	csvArray := generateCostCSVArray(services)
	writeCSV(csvArray, "test.csv")
	t.Run("TestFileExistance", func(t *testing.T) {
		if _, err := os.Stat("test.csv"); os.IsNotExist(err) {
			t.Errorf("Expected file to exist")
		}
	})

	t.Run("TestFileContent", func(t *testing.T) {
		file, err := os.Open("test.csv")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		defer file.Close()
		fileContent := make([]byte, 1024)
		_, err = file.Read(fileContent)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if strings.Split(string(fileContent), "\n")[0] != "Service,2019-01-01,2019-02-01,2019-03-01,2019-04-01,2019-05-01" {
			t.Errorf("Expected file content to be correct, got %s", strings.Split(string(fileContent), "\n")[0])
		}
	})
	//Delete file
	os.Remove("test.csv")
}

func TestInitCSV(t *testing.T) {
	var services Services
	services = append(services, service)

	t.Run("TestFileExistance", func(t *testing.T) {
		flag.Set("csv", "true")
		services.initCSV()
		if _, err := os.Stat("reports.csv"); os.IsNotExist(err) {
			t.Errorf("Expected file to exist")
		}
		os.Remove("evolution.csv")
		os.Remove("reports.csv")
	})

	t.Run("TestNoFile", func(t *testing.T) {
		flag.Set("csv", "false")
		services.initCSV()
		if _, err := os.Stat("reports.csv"); err == nil {
			t.Errorf("Expected file to not exist")
		}
		os.Remove("evolution.csv")
		os.Remove("reports.csv")
	})

	t.Run("TestEvolutionOnlyFlag", func(t *testing.T) {
		flag.Set("csv", "false")
		flag.Set("evolution", "true")
		services.initCSV()
		if _, err := os.Stat("evolution.csv"); err == nil {
			t.Errorf("Expected file to not exist")
		}
		if _, err := os.Stat("reports.csv"); err == nil {
			t.Errorf("Expected file to not exist")
		}
		os.Remove("evolution.csv")
		os.Remove("reports.csv")
	})

	t.Run("TestEvolution", func(t *testing.T) {
		flag.Set("csv", "true")
		flag.Set("evolution", "true")
		services.initCSV()
		if _, err := os.Stat("evolution.csv"); os.IsNotExist(err) {
			t.Errorf("Expected file to exist")
		}
		os.Remove("evolution.csv")
		os.Remove("reports.csv")
	})

}
