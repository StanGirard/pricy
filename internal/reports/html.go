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

	"github.com/stangirard/pricy/internal/format"
)

//go:embed templates/content.tmpl
var contentTmpl string

//go:embed templates/footer.tmpl
var footerTmpl string

//go:embed templates/header.tmpl
var headerTmpl string

//go:embed templates/page.tmpl
var pageTmpl string

//go

var (
	html = flag.Bool("html", false, "print in html format")
)

func getPriceForDateService(service format.Service, date format.DateInterval) float64 {

	stringValue := fmt.Sprintf("%.2f", service.DatePrice[date])
	fl, _ := strconv.ParseFloat(stringValue, 64)
	return fl
}

func (services Services) generateHTML() string {
	dates := format.SortDates(format.FindDatesIntervals(services))
	// Create a template containing contentTmpl and footerTmpl and headerTmpl and pageTmpl
	tmpl, err := template.New("content").Funcs(template.FuncMap{"getPrice": getPriceForDateService}).Parse(contentTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.New("footer").Parse(footerTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.New("header").Parse(headerTmpl)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.New("page").Parse(pageTmpl)

	if err != nil {
		log.Fatal(err)
	}
	// Create a buffer to hold the generated HTML
	var processed bytes.Buffer
	// Pass dates and services to the template.
	tmpl.Execute(&processed, struct {
		Title    string
		Dates    []format.DateInterval
		Services []format.Service
	}{
		Title:    "Pricy",
		Dates:    dates,
		Services: services,
	})
	fmt.Println(processed.String())
	return processed.String()

}

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

func (services Services) initHTML() {
	if *html {
		html := services.generateHTML()
		writeToFile(html)
	}
}
