package gsheet

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/sheets/v4"
)

// Creates the spreadsheet configuration
// Mostly used for setting the title of the spreadsheet that will be created
func createSpreadSheetConfig() *sheets.Spreadsheet {
	// Generate the random time
	var timeNowString = time.Now().Format("2006-01-02-15-04-05")
	title := "Pricy-Report-" + timeNowString

	// Create the spreadsheet configuration
	// Can be customed further to set the title, etc.
	rb := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
	}
	return rb
}

// Makes a BatchUpdateRequest to add a sheet to the spreadsheet
// Takes the spreadsheet service, the spreadsheet id, the sheet title and the sheet index
func addSheet(service *sheets.Service, spreadsheetId, name string, idSheet int64) {
	_, error := service.Spreadsheets.BatchUpdate(*spreadsheet, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: &sheets.SheetProperties{
						Title:   name,
						SheetId: idSheet,
					},
				},
			},
		},
	}).Do()

	if error != nil {
		log.Fatalf("Unable to add sheet: %v", error)
	}
	fmt.Println("Sheet added: ", name)
}

// Adds a basic Chart to the spreadsheet
//
// Takes the spreadsheet service, the spreadsheet id, the sheet title and the sheet index to which the chart will be added and the specs of the chart
func addChart(service *sheets.Service, spreadsheetId string, idSheetToAddTo int64, specs *sheets.BasicChartSpec) {
	_, error := service.Spreadsheets.BatchUpdate(*spreadsheet, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddChart: &sheets.AddChartRequest{
					Chart: &sheets.EmbeddedChart{
						Position: &sheets.EmbeddedObjectPosition{
							NewSheet: false,
							OverlayPosition: &sheets.OverlayPosition{
								AnchorCell: &sheets.GridCoordinate{
									SheetId:     idSheetToAddTo,
									RowIndex:    0,
									ColumnIndex: 0,
								},
								WidthPixels:  1200,
								HeightPixels: 600,
							},
						},
						Spec: &sheets.ChartSpec{
							BasicChart: specs,
						},
					},
				},
			},
		},
	}).Do()

	if error != nil {
		log.Fatalf("Unable to add sheet: %v", error)
	}
	fmt.Println("Chart added")
}

// Delete a sheet
// Takes the spreadsheet service, the spreadsheet id and the sheet id to delete
func deleteSheet(service *sheets.Service, spreadsheetId string, sheetId int64) {
	_, error := service.Spreadsheets.BatchUpdate(*spreadsheet, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				DeleteSheet: &sheets.DeleteSheetRequest{
					SheetId: sheetId,
				},
			},
		},
	}).Do()

	if error != nil {
		fmt.Printf("Sheet not deleted because it does not exist (Id:%v)\n", sheetId)
	}
	fmt.Println("Default sheet deleted")
}

// Populate the spreadsheet with the data
// The Data must be an array of arrays
func writeData(context context.Context, service *sheets.Service, spreadsheetId string, data [][]interface{}) {
	valuesRanges := make([]*sheets.ValueRange, len(data))
	valuesRanges = append(valuesRanges, &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         data,
		Range:          "A1:ZZ",
	},
	)
	rb2 := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
		Data:             valuesRanges,
	}

	_, err := service.Spreadsheets.Values.BatchUpdate(*spreadsheet, rb2).Context(context).Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sheet populated")
}

// Create a new empty spreadshit
func createSpreadsheet(service *sheets.Service, config *sheets.Spreadsheet, ctx context.Context) string {
	resp, err := service.Spreadsheets.Create(config).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to create spreadsheet: %v", err)
	}
	return resp.SpreadsheetId
}
