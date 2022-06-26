package main

//
//import (
//	"context"
//	"fmt"
//	"github.com/joho/godotenv"
//	"google-sheet-sample/infra"
//	"google.golang.org/api/option"
//	"google.golang.org/api/sheets/v4"
//	"gopkg.in/yaml.v3"
//	"io/ioutil"
//	"log"
//)
//
//func main() {
//	// 設定ファイル読みこみ
//	data, err := ioutil.ReadFile("config.yaml")
//	if err != nil {
//		log.Fatalf("failed to read config.yaml: %v", err)
//	}
//	config := infra.Config{}
//	err = yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatalf("failed to struct config: %v", err)
//	}
//
//	err = godotenv.Load()
//	if err != nil {
//		log.Fatalf("Error loading .env file: %v", err)
//	}
//	credentialFileName := "credential.json"
//	//spreadsheetId := os.Getenv("SPREADSHEET_ID")
//	//sheetId, err := strconv.Atoi(os.Getenv("SHEET_ID"))
//	//if err != nil {
//	//	log.Fatalf("failed to cast string to integer: %v", err)
//	//}
//
//	credential := option.WithCredentialsFile(credentialFileName)
//	service, err := sheets.NewService(context.TODO(), credential)
//	if err != nil {
//		log.Fatalf("failed to create google spread sheet service: %v", err)
//	}
//
//	rangeVal := "シート4"
//	ctx := context.Background()
//	resp, err := service.Spreadsheets.Values.Get("19shjxKXH1apr51YCluuJr4dKxSviTITPA5dd8cWgNLA", rangeVal).ValueRenderOption("FORMATTED_VALUE").Context(ctx).Do()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("%#v\n", resp)
//}
