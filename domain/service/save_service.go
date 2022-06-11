package service

import (
	"google-sheet-sample/domain/client"
	"log"
	"strconv"
	"time"
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
	storeInfo, err := s.storeClient.Get()
	if err != nil {
		log.Println("failed to get store info")
		return err
	}

	// 3. 1. 2. で取得したデータを突合して更新データを生成
	result := make([]string, len(storeInfo)+1)
	result = append(result, time.Now().Format("2006/Jan/02"))
	for _, e := range storeInfo {
		if v, ok := favoriteMap[e.GetId()]; ok {
			result = append(result, strconv.Itoa(v))
		} else {
			result = append(result, "0")
		}
	}

	// 4. データを保存
	err = s.dataManager.Save(result)
	if err != nil {
		log.Println("failed to update favorite data")
		return err
	}

	log.Println("data save success")
	return nil
}
