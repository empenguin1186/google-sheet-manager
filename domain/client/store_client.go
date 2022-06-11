package client

import "google-sheet-sample/domain/model"

type StoreClient interface {
	Get() ([]*model.StoreInfo, error)
}
