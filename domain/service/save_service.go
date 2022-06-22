package service

import (
	"google-sheet-sample/domain/client"
	"google-sheet-sample/domain/model"
	"log"
)

type SaveService struct {
	favoriteClient client.FavoriteClient
	storeClient    client.StoreClient
	dataManager    client.DataManager
}

func NewSaveService(favoriteClient client.FavoriteClient, storeClient client.StoreClient, dataManager client.DataManager) *SaveService {
	return &SaveService{favoriteClient: favoriteClient, storeClient: storeClient, dataManager: dataManager}
}

func (s *SaveService) Save() error {
	// 1. 各店舗のお気に入り数を取得
	favoriteMap, err := s.favoriteClient.Get()
	if err != nil {
		log.Println("failed to get favorite data")
		return err
	}

	// 2. 店舗情報取得
	storeInfoMap, max, err := s.storeClient.Get()
	if err != nil {
		log.Println("failed to get store info")
		return err
	}

	// 3. 1. 2. で取得したデータを突合して更新データを生成
	var storeIds []int
	var storeNames []string
	var favorites []int

	// 1~(店舗IDの最大値)の間で下記の処理を繰り返す
	// 登録店舗に更新があった場合にデータの不整合が発生しないようにするため連番で管理
	// 店舗情報が削除され葉抜きのIDになった場合にデータの不整合が発生するのでそれを回避
	for i := 0; i < max; i++ {
		storeId := i + 1
		storeIds = append(storeIds, storeId)

		if v, ok := storeInfoMap[storeId]; ok {
			storeNames = append(storeNames, v)
		} else {
			storeNames = append(storeNames, "-")
		}

		if v, ok := favoriteMap[storeId]; ok {
			favorites = append(favorites, v)
		} else {
			favorites = append(favorites, 0)
		}
	}

	saveData, err := model.NewSaveData(storeIds, storeNames, favorites)
	if err != nil {
		log.Println("failed to construct save data")
		return err
	}

	// 4. データを保存
	err = s.dataManager.Save(&saveData)
	if err != nil {
		log.Println("failed to update favorite data")
		return err
	}

	log.Println("data save success")
	return nil
}
