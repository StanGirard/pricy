package reports

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/stangirard/pricy/internal/format"
)

//go:embed templates/cost.tmpl
var costTmpl string

//go:embed templates/evolution.tmpl
var evolutionTmpl string

//go:embed templates/footer.tmpl
var footerTmpl string

//go:embed templates/header.tmpl
var headerTmpl string

//go:embed templates/page.tmpl
var pageTmpl string

//go:embed templates/style.tmpl
var styleTmpl string

//go

var (
	html = flag.Bool("html", false, "print in html format") // Flag to print an html report
)

// Gets the price of the service given a dateinterval and formats it with 2 decimals
func getPriceForDateService(service format.Service, date format.DateInterval) float64 {

	stringValue := fmt.Sprintf("%.2f", service.DatePrice[date])
	fl, _ := strconv.ParseFloat(stringValue, 64)
	return fl
}

// Gets the evolution of the service given a dateinterval and formats it with 2 decimals
func getEvolutionForDateService(service format.Service, date format.DateInterval) float64 {

	stringValue := fmt.Sprintf("%.2f", service.PriceEvolution[date])
	fl, _ := strconv.ParseFloat(stringValue, 64)
	return fl
}

// Pretty print the date as DD/MM in order to take less space in the html report
func niceDate(date string) string {
	t, _ := time.Parse("2006-01-02", date)
	return t.Format("02/01")

}

// Merges all tpl file into one template (Files are in templates/)
func createTemplate() (*template.Template, error) {
	tmpl, err := template.New("cost").Funcs(template.FuncMap{"getPrice": getPriceForDateService, "niceDate": niceDate}).Parse(costTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.New("footer").Parse(footerTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.New("evolution").Funcs(template.FuncMap{"getEvolution": getEvolutionForDateService}).Parse(evolutionTmpl)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = tmpl.New("style").Parse(styleTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.New("header").Parse(headerTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.New("page").Parse(pageTmpl)
	return tmpl, err
}

// Generate the html given the services
func (services Services) generateHTML() string {
	dates := format.SortDates(format.FindDatesIntervals(services))

	// Create a template containing contentTmpl and footerTmpl and headerTmpl and pageTmpl
	tmpl, err := createTemplate()
	if err != nil {
		log.Fatal(err)
	}
	// Create a buffer to hold the generated HTML
	var processed bytes.Buffer
	// Pass dates and services to the template.
	tmpl.Execute(&processed, struct {
		Title     string
		Dates     []format.DateInterval
		Services  []format.Service
		Evolution bool
	}{
		Title:     "Cost of services",
		Dates:     dates,
		Services:  services,
		Evolution: *evolution,
	})
	return processed.String()

}

// Write the html into a file
func writeToFile(html string) {
	f, err := os.Create("pricy.html")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(html)
	if err != nil {
		log.Fatal(err)
	}

}

// Logic around flags of the html reports
func (services Services) initHTML() {
	flag.Parse()
	if *html {
		html := services.generateHTML()
		writeToFile(html)
	}
}
