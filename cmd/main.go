package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	credentialFileName := os.Getenv("CREDENTIAL_FILENAME")
	spreadsheetId := os.Getenv("SHEET_ID")

	credential := option.WithCredentialsFile(credentialFileName)
	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err)
	}
	readRange := "シート1!A1:B3"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalln(err)
	}
	if len(resp.Values) == 0 {
		log.Fatalln("data not found")
	}
	for _, row := range resp.Values {
		fmt.Printf("%s, %s\n", row[0], row[1])
	}
}
