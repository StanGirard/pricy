package gsheet

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	spreadsheet = flag.String("spreadsheet", "", "The ID of the spreadsheet to write to.")
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	authCode = httpServerParseResponse()

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func httpServerParseResponse() string {
	// Create a new HTTP server that waits for a request, then parses it and sends a response.
	// Request is http://localhost:6767/?state=state-token&code=4/0AdQt8qgrpmO_twv6d32zwSKxlSWE6Z4tHwcM6HHlyqWu4UNRJmC6faLYSBODOZ5H0xOlWw&scope=https://www.googleapis.com/auth/spreadsheets

	ctx, cancel := context.WithCancel(context.Background())
	var response string
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request.
		err := r.ParseForm()
		if err != nil {
			log.Print(err)
		}
		// Print the request parameters.
		fmt.Println(r.Form)
		// Write to w "Token has been received and saved into token.json"
		fmt.Fprint(w, "Token has been received and saved into token.json, you can close this window.")

		response = r.Form.Get("code")
		cancel()
	})
	srv := &http.Server{Addr: ":6767"}

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	// gracefully shutdown the server:
	// waiting indefinitely for connections to return to idle and then shut down.
	err := srv.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}

	log.Println("done.")
	return response

}

func Execute(value [][]string) {
	flag.Parse()
	old := value
	new := make([][]interface{}, len(old))
	for i, v := range old {
		new[i] = make([]interface{}, len(v))
		for j, v2 := range v {
			new[i][j] = strings.Replace(v2, ".", ",", -1)
		}
	}

	ctx := context.Background()
	//Os get env variable  GOOGLE_APPLICATION_CREDENTIALS
	GOOGLE_APPLICATION_CREDENTIALS := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	b, err := ioutil.ReadFile(GOOGLE_APPLICATION_CREDENTIALS)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Create a new spreadsheet.
	var timeNowString = time.Now().Format("2006-01-02-15-04-05")
	title := "Pricy-Report-" + timeNowString
	rb := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
	}

	var BasicChartSeries []*sheets.BasicChartSeries
	var HistogramSeries []*sheets.HistogramSeries
	for i, _ := range new {
		if i == 0 {
			continue
		}
		BasicChartSeries = append(BasicChartSeries, &sheets.BasicChartSeries{
			Series: &sheets.ChartData{
				SourceRange: &sheets.ChartSourceRange{
					Sources: []*sheets.GridRange{
						{ //A2:O2
							SheetId:          1023,
							StartRowIndex:    int64(i),
							EndRowIndex:      int64(i + 1),
							StartColumnIndex: 0,
							EndColumnIndex:   int64(len(new[i])),
						},
					},
				},
			},
		},
		)
		HistogramSeries = append(HistogramSeries, &sheets.HistogramSeries{
			Data: &sheets.ChartData{
				SourceRange: &sheets.ChartSourceRange{
					Sources: []*sheets.GridRange{
						{ //A2:O2
							SheetId:          1023,
							StartRowIndex:    int64(i),
							EndRowIndex:      int64(i + 1),
							StartColumnIndex: 0,
							EndColumnIndex:   int64(len(new[i])),
						},
					},
				},
			},
		})
	}

	if *spreadsheet == "" {
		resp, err := srv.Spreadsheets.Create(rb).Context(ctx).Do()
		if err != nil {
			log.Fatalf("Unable to create spreadsheet: %v", err)
		}
		*spreadsheet = resp.SpreadsheetId
		_, error := srv.Spreadsheets.BatchUpdate(*spreadsheet, &sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				{
					AddSheet: &sheets.AddSheetRequest{
						Properties: &sheets.SheetProperties{
							Title:   "Reports",
							SheetId: 1023,
						},
					},
				},
				{
					AddSheet: &sheets.AddSheetRequest{
						Properties: &sheets.SheetProperties{
							Title:   "Charts",
							SheetId: 1024,
						},
					},
				},
				// {
				// 	AddChart: &sheets.AddChartRequest{
				// 		Chart: &sheets.EmbeddedChart{
				// 			Position: &sheets.EmbeddedObjectPosition{
				// 				NewSheet: false,
				// 				OverlayPosition: &sheets.OverlayPosition{
				// 					AnchorCell: &sheets.GridCoordinate{
				// 						SheetId:     1024,
				// 						RowIndex:    1,
				// 						ColumnIndex: 1,
				// 					},
				// 				},
				// 			},
				// 			Spec: &sheets.ChartSpec{
				// 				HistogramChart: &sheets.HistogramChartSpec{
				// 					Series:     HistogramSeries,
				// 					BucketSize: float64(len(new)),
				// 				},
				// 			},
				// 		},
				// 	},
				// },
				{
					AddChart: &sheets.AddChartRequest{
						Chart: &sheets.EmbeddedChart{
							Position: &sheets.EmbeddedObjectPosition{
								NewSheet: false,
								OverlayPosition: &sheets.OverlayPosition{
									AnchorCell: &sheets.GridCoordinate{
										SheetId:     1024,
										RowIndex:    0,
										ColumnIndex: 0,
									},
									WidthPixels:  1200,
									HeightPixels: 600,
								},
							},
							Spec: &sheets.ChartSpec{
								BasicChart: &sheets.BasicChartSpec{
									HeaderCount:    1,
									Series:         BasicChartSeries,
									LegendPosition: "BOTTOM_LEGEND",
									ChartType:      "COLUMN",
									StackedType:    "STACKED",
									Domains: []*sheets.BasicChartDomain{
										{
											Domain: &sheets.ChartData{
												SourceRange: &sheets.ChartSourceRange{
													Sources: []*sheets.GridRange{
														{
															SheetId:          1023,
															StartRowIndex:    0,
															EndRowIndex:      1,
															StartColumnIndex: 0,
															EndColumnIndex:   int64(len(new[0])),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					DeleteSheet: &sheets.DeleteSheetRequest{
						SheetId: 0,
					},
				},
			},
		}).Do()
		if error != nil {
			log.Fatalf("Unable to create sheet: %v", error)
		}
	}

	// Write reports.csv to the new spreadsheet.

	// Write the value variable to the new spreadsheet.
	valuesRanges := make([]*sheets.ValueRange, len(new))
	valuesRanges = append(valuesRanges, &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         new,
		Range:          "A1:ZZ",
	},
	)
	rb2 := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
		Data:             valuesRanges,
	}

	_, err = srv.Spreadsheets.Values.BatchUpdate(*spreadsheet, rb2).Context(ctx).Do()

	if err != nil {
		log.Fatal(err)
	}

	// Print the spreadsheet url
	fmt.Printf("Spreadsheet URL: https://docs.google.com/spreadsheets/d/%s\n", *spreadsheet)
}
