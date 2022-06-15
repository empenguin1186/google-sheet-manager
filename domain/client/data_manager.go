package client

import "google-sheet-sample/domain/model"

type DataManager interface {
	Save(data *model.SaveData) error
}
