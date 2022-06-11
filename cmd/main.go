package main

//
//import (
//	"context"
//	"fmt"
//	"github.com/joho/godotenv"
//	"google.golang.org/api/option"
//	"google.golang.org/api/sheets/v4"
//	"log"
//	"os"
//)
//
//func main() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	credentialFileName := os.Getenv("CREDENTIAL_FILENAME")
//	spreadsheetId := os.Getenv("SPREADSHEET_ID")
//
//	credential := option.WithCredentialsFile(credentialFileName)
//	srv, err := sheets.NewService(context.TODO(), credential)
//	if err != nil {
//		log.Fatal(err)
//	}
//	readRange := "シート1"
//	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	if len(resp.Values) == 0 {
//		log.Fatalln("data not found")
//	}
//	for _, row := range resp.Values {
//		for _, col := range row {
//			fmt.Printf("%s, ", col)
//		}
//		fmt.Println("")
//	}
//
//	data := []string{"2022/06/10", "aaa", "bbb", "ccc"}
//	vals := make([]*sheets.CellData, 0, len(data))
//	for i := range data {
//		vals = append(vals, &sheets.CellData{
//			UserEnteredValue: &sheets.ExtendedValue{
//				StringValue: &data[i],
//			},
//		})
//	}
//
//	req := []*sheets.Request{
//		{
//			AppendCells: &sheets.AppendCellsRequest{
//				SheetId: 0,
//				Fields:  "*",
//				Rows: []*sheets.RowData{
//					{
//						Values: vals,
//					},
//				},
//			},
//		},
//	}
//
//	_, err = srv.Spreadsheets.BatchUpdate(spreadsheetId, &sheets.BatchUpdateSpreadsheetRequest{
//		IncludeSpreadsheetInResponse: false,
//		Requests:                     req,
//		ResponseIncludeGridData:      false,
//	}).Do()
//	if err != nil {
//		log.Fatal(err)
//	}
//}
