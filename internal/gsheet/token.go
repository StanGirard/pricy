package gsheet

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
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
