package gsheet

import (
	"flag"
	"fmt"

	"github.com/stangirard/pricy/internal/helpers"
)

var (
	spreadsheet = flag.String("spreadsheet", "", "The ID of the spreadsheet to write to.")
)

// Execute is the main function of the gsheet package.
func Execute(value [][]string) {
	flag.Parse()

	dataSpreadsheet := helpers.ConvertStringToInterface(value)
	srv, ctx := NewSheetService()

	// Create a new spreadsheet.
	rb := createSpreadSheetConfig()

	BasicChartSeries := createBasicChartSeries(dataSpreadsheet)

	if *spreadsheet == "" {
		*spreadsheet = createSpreadsheet(srv, rb, ctx)
		// Add a sheet
		addSheet(srv, *spreadsheet, "Reports", 1023)
		deleteSheet(srv, *spreadsheet, 0)

	}

	//Delete the Chart sheet if it exists
	deleteSheet(srv, *spreadsheet, 1024)
	addSheet(srv, *spreadsheet, "Charts", 1024)

	// Write the value variable to the new spreadsheet.
	writeData(ctx, srv, *spreadsheet, dataSpreadsheet)

	// Create BasicChart Series
	chartSpecs := createBasicChartSpec("BOTTOM_LEGEND", "COLUMN", "STACKED", 1023, dataSpreadsheet, BasicChartSeries)
	addChart(srv, *spreadsheet, 1024, chartSpecs)

	// Print the spreadsheet url
	fmt.Printf("Spreadsheet URL: https://docs.google.com/spreadsheets/d/%s\n", *spreadsheet)
}
