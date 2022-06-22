package main

import (
	"github.com/joho/godotenv"
	"google-sheet-sample/domain/service"
	"google-sheet-sample/infra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	// 設定ファイル読みこみ
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("failed to read config.yaml: %v", err)
	}
	config := infra.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("failed to struct config: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	credentialFileName := "empentech.json"
	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	sheetId, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	if err != nil {
		log.Fatalf("failed to cast string to integer: %v", err)
	}
	idToken := os.Getenv("ID_TOKEN")

	// 各種構造体構築
	yakkyubinClient := infra.NewYakkyubinClient(&config.Yakyuubin)
	googleSpreadSheetManager, err := infra.NewGoogleSpreadSheetService(credentialFileName, spreadsheetId, sheetId)
	if err != nil {
		log.Fatalf("failed to struct GoogleSpreadSheetManager: %v", err)
	}
	favoriteClientImpl := infra.NewFavoriteClientImpl(idToken, &config.Favorite)

	saveService := service.NewSaveService(favoriteClientImpl, yakkyubinClient, googleSpreadSheetManager)

	// 処理実行
	err = saveService.Save()
	if err != nil {
		log.Fatalf("failed to save data: %v", err)
	}
}
