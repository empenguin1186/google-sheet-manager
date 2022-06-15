package infra

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

type GoogleSpreadSheetService struct {
	spreadSheetService *sheets.Service
	spreadSheetId      string
	sheetId            int
}

func NewGoogleSpreadSheetService(credentialFileName string, spreadSheetId string, sheetId int) (*GoogleSpreadSheetService, error) {
	credential := option.WithCredentialsFile(credentialFileName)
	service, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err)
		return &GoogleSpreadSheetService{}, err
	}
	return &GoogleSpreadSheetService{spreadSheetService: service, spreadSheetId: spreadSheetId, sheetId: sheetId}, nil
}

func (g *GoogleSpreadSheetService) Save(data []string) error {

	// セルデータに変換
	var values []*sheets.CellData
	for i, _ := range data {
		values = append(values, &sheets.CellData{
			UserEnteredValue: &sheets.ExtendedValue{
				StringValue: &data[i],
			},
		})
	}

	// 行データに変換
	rowData := []*sheets.RowData{
		{
			Values: values,
		},
	}

	// リクエスト作成
	// 登録店舗が更新される可能性を考慮しSpreadSheetの1行目(RowIndex=0)に登録されている店舗IDは毎回更新する
	updateCellsRequest1 := sheets.UpdateCellsRequest{
		Fields: "*",
		Rows:   rowData,
		Start: &sheets.GridCoordinate{
			SheetId:     0,
			RowIndex:    0,
			ColumnIndex: 0,
		},
	}

	// 登録店舗が更新される可能性を考慮しSpreadSheetの2行目(RowIndex=1)に登録されている店舗名は毎回更新する
	updateCellsRequest2 := sheets.UpdateCellsRequest{
		Fields: "*",
		Rows:   rowData,
		Start: &sheets.GridCoordinate{
			SheetId:     0,
			RowIndex:    0,
			ColumnIndex: 0,
		},
	}

	// 今日の分のお気に入り登録数を行の末尾に追加するためのリクエストを構築
	appendCellsRequest := sheets.AppendCellsRequest{
		Fields:  "*",
		Rows:    rowData,
		SheetId: 0,
	}

	requests := []*sheets.Request{
		{
			UpdateCells: &updateCellsRequest1,
		},
		{
			UpdateCells: &updateCellsRequest2,
		},
		{
			AppendCells: &appendCellsRequest,
		},
	}
	batchRequest := &sheets.BatchUpdateSpreadsheetRequest{
		IncludeSpreadsheetInResponse: true,
		Requests:                     requests,
	}

	ctx := context.Background()
	_, err := g.spreadSheetService.Spreadsheets.BatchUpdate(g.spreadSheetId, batchRequest).Context(ctx).Do()

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
