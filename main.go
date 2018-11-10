package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aphill70/sheet-rotation/gifts"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
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
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

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

func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1y3ySYxxxsmLRSZBKJz0CFk9KCdvm8H38pjMcH_Uzixk"
	readRange := "Rotation!A1:Z"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var rotations []*gifts.Rotation
	var currentRotation *gifts.Rotation
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		columnMap := make(map[int]string)

		for rowIndex, row := range resp.Values {
			if len(row) == 0 || rowIndex == 0 {
				fmt.Println("found empty row initializing new rotation")
				if currentRotation != nil {
					rotations = append(rotations, currentRotation)
				}
				currentRotation, _ = gifts.NewRotation()
			}

			for column, cell := range row {
				if rowIndex == 0 {
					// handle headers
					columnMap[column] = fmt.Sprintf("%s", cell)
				} else {
					if columnMap[column] == "Name" {
						currentRotation.AddRecipient(fmt.Sprintf("%s", cell))
					} else {
						giver := fmt.Sprintf("%s", cell)
						if !strings.Contains(giver, "X") {
							currentRotation.AddGiver(giver, columnMap[column])
						}
					}

					fmt.Printf(" %s:%s ", columnMap[column], cell)
				}
			}
			fmt.Println("")
		}
		if len(rotations) > 0 {
			err := rotations[0].GetEligibleGivers("Pop")
			if err != nil {
				os.Exit(1)
			}
		}
	}
}
