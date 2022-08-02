package gsheet

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewSheetService() (*sheets.Service, context.Context) {

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

	return srv, ctx
}
