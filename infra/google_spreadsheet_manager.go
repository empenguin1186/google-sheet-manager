package infra

import (
	"context"
	"google-sheet-sample/domain/model"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"time"
)

type GoogleSpreadSheetService struct {
	spreadSheetService *sheets.Service
	spreadSheetId      string
	sheetId            int64
}

func NewGoogleSpreadSheetService(credentialFileName string, spreadSheetId string, sheetId int64) (*GoogleSpreadSheetService, error) {
	credential := option.WithCredentialsFile(credentialFileName)
	service, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err)
		return &GoogleSpreadSheetService{}, err
	}
	return &GoogleSpreadSheetService{spreadSheetService: service, spreadSheetId: spreadSheetId, sheetId: sheetId}, nil
}

func (g *GoogleSpreadSheetService) Save(data *model.SaveData) error {

	// 行データに変換
	storeIdsRowData := g.makeNumberRowData("店舗ID", data.GetStoreIds())
	storeNamesRowData := g.makeStringRowData("店舗名", data.GetStoreNames())
	favoritesRowData := g.makeNumberRowData(time.Now().Format("2006/01/02"), data.GetFavorites())

	// リクエスト作成
	// 登録店舗が更新される可能性を考慮しSpreadSheetの1行目(RowIndex=0)に登録されている店舗IDは毎回更新する
	updateCellsRequest1 := sheets.UpdateCellsRequest{
		Fields: "*",
		Rows:   storeIdsRowData,
		Start: &sheets.GridCoordinate{
			SheetId:     g.sheetId,
			RowIndex:    0,
			ColumnIndex: 0,
		},
	}

	// 登録店舗が更新される可能性を考慮しSpreadSheetの2行目(RowIndex=1)に登録されている店舗名は毎回更新する
	updateCellsRequest2 := sheets.UpdateCellsRequest{
		Fields: "*",
		Rows:   storeNamesRowData,
		Start: &sheets.GridCoordinate{
			SheetId:     g.sheetId,
			RowIndex:    1,
			ColumnIndex: 0,
		},
	}

	// 今日の分のお気に入り登録数を行の末尾に追加するためのリクエストを構築
	appendCellsRequest := sheets.AppendCellsRequest{
		Fields:  "*",
		Rows:    favoritesRowData,
		SheetId: g.sheetId,
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

func (g *GoogleSpreadSheetService) makeNumberRowData(item string, data []int) []*sheets.RowData {
	// セルデータに変換
	var values []*sheets.CellData

	// 最初に項目を追加
	values = append(values, &sheets.CellData{
		UserEnteredValue: &sheets.ExtendedValue{
			StringValue: &item,
		},
	})

	// 実際のデータを追加
	for i, _ := range data {
		value := float64(data[i])
		values = append(values, &sheets.CellData{
			UserEnteredValue: &sheets.ExtendedValue{
				NumberValue: &value,
			},
		})
	}

	// 行データに変換
	rowData := []*sheets.RowData{
		{
			Values: values,
		},
	}

	return rowData
}

func (g *GoogleSpreadSheetService) makeStringRowData(item string, data []string) []*sheets.RowData {
	// セルデータに変換
	var values []*sheets.CellData

	// 最初に項目を追加
	values = append(values, &sheets.CellData{
		UserEnteredValue: &sheets.ExtendedValue{
			StringValue: &item,
		},
	})

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

	return rowData
}
