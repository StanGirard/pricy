package reports

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/stangirard/pricy/internal/format"
)

func TestGetPriceForDateService(t *testing.T) {
	t.Run("GetPriceForDateService", func(t *testing.T) {
		value := getPriceForDateService(service, format.DateInterval{Start: "2019-01-01", End: "2019-01-31"})
		if value != 0.1 {
			t.Errorf("Expected 0.1, got %f", value)
		}
	})
}

func TestGetEvolutionForDateService(t *testing.T) {
	t.Run("GetEvolutionForDateService", func(t *testing.T) {
		value := getEvolutionForDateService(service, format.DateInterval{Start: "2019-01-01", End: "2019-01-31"})
		if value != 0.0 {
			t.Errorf("Expected 0.0, got %f", value)
		}
	})
}

func TestNiceDate(t *testing.T) {
	t.Run("NiceDate", func(t *testing.T) {
		value := niceDate("2019-01-01")
		if value != "01/01" {
			t.Errorf("Expected 01/01, got %s", value)
		}
	})
}

func TestCreateTemplate(t *testing.T) {
	t.Run("CreateTemplate", func(t *testing.T) {
		template, err := createTemplate()

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
		if template == nil {
			t.Errorf("Expected not nil, got nil")
		}
	})
}

func TestGenerateHTML(t *testing.T) {
	var services Services
	services = append(services, service)
	t.Run("GenerateHTML", func(t *testing.T) {
		html := services.generateHTML()
		if html == "" {
			t.Errorf("Expected not nil, got nil")
		}
	})
	t.Run("CheckContent", func(t *testing.T) {
		html := services.generateHTML()
		if !strings.Contains(html, "Cost of services") {
			t.Errorf("Expected to contain Cost of services, got %s", html)
		}
	})
}

func TestWriteToFile(t *testing.T) {
	var services Services
	services = append(services, service)
	t.Run("writeToFile", func(t *testing.T) {
		html := services.generateHTML()
		writeToFile(html)
		// Check if file exists
		if _, err := os.Stat("./pricy.html"); os.IsNotExist(err) {
			t.Errorf("Expected file to exist, got %s", err)
		}
		//Delete file
		os.Remove("./pricy.html")
	})
}

func TestInitHTML(t *testing.T) {
	var services Services
	services = append(services, service)

	t.Run("InitHTML", func(t *testing.T) {
		flag.Set("html", "true")
		services.initHTML()
		// Check if file exists
		if _, err := os.Stat("./pricy.html"); os.IsNotExist(err) {
			t.Errorf("Expected file to exist, got %s", err)
		}
		//Delete file
		os.Remove("./pricy.html")
	})
	t.Run("InitHTML", func(t *testing.T) {
		flag.Set("html", "false")
		services.initHTML()
		// Check if file does not exist
		if _, err := os.Stat("./pricy.html"); err == nil {
			t.Errorf("Expected file to not exist, got %s", err)
		}
	})
}
