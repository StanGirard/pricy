package reports

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/stangirard/pricy/internal/format"
)

var (
	html = flag.Bool("html", false, "print in html format")
)

func (services Services) generateHTML() string {
	var html string
	html += "<html><head><title>Pricy</title></head><body>"
	html += "<h1>Pricy</h1>"
	html += "<table>"
	html += "<tr><th>Service</th>"
	dates := format.SortDates(format.FindDatesIntervals(services))
	for _, dates := range dates {
		html += fmt.Sprint("<th>", dates.Start, "</th>")
	}
	html += "</tr>"
	for _, service := range services {
		html += "<tr>"
		html += fmt.Sprint("<td><b>", service.Name, "</b></td>")
		for _, date := range dates {
			// Maximum 2 decimal places for prices
			html += fmt.Sprint("<td>", fmt.Sprintf("%.2f", service.DatePrice[date]), "</td>")
		}
		html += "</tr>"
	}

	html += "</table>"
	html += "</body></html>"
	return html

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
