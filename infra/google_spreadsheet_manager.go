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
	vals := make([]*sheets.CellData, 0, len(data))
	for i := range data {
		vals = append(vals, &sheets.CellData{
			UserEnteredValue: &sheets.ExtendedValue{
				StringValue: &data[i],
			},
		})
	}

	request := []*sheets.Request{
		{
			AppendCells: &sheets.AppendCellsRequest{
				SheetId: 0,
				Fields:  "*",
				Rows: []*sheets.RowData{
					{
						Values: vals,
					},
				},
			},
		},
	}

	_, err := g.spreadSheetService.Spreadsheets.BatchUpdate(g.spreadSheetId, &sheets.BatchUpdateSpreadsheetRequest{
		IncludeSpreadsheetInResponse: false,
		Requests:                     request,
		ResponseIncludeGridData:      false,
	}).Do()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
