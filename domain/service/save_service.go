package service

import (
	"google-sheet-sample/domain/client"
	"google-sheet-sample/domain/model"
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
	storeInfoMap, max, err := s.storeClient.Get()
	if err != nil {
		log.Println("failed to get store info")
		return err
	}

	// 3. 1. 2. で取得したデータを突合して更新データを生成
	var storeIds []string
	var storeNames []string
	var favorites []string
	storeIds = append(storeIds, "店舗ID")
	storeNames = append(storeNames, "店舗名")
	favorites = append(favorites, time.Now().Format("2006/01/02"))

	for i := 1; i <= max; i++ {
		storeIds = append(storeIds, strconv.Itoa(i))

		if v, ok := storeInfoMap[i]; ok {
			storeNames = append(storeNames, v)
		} else {
			storeNames = append(storeNames, "-")
		}

		if v, ok := favoriteMap[i]; ok {
			favorites = append(favorites, strconv.Itoa(v))
		} else {
			favorites = append(favorites, "0")
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

	// 店舗IDの1~最大値の配列を作成
	// 登録店舗に更新があった場合にデータの不整合が発生しないようにするため連番で管理
	// 店舗情報が削除され葉抜きのIDになった場合にデータの操作が面倒になる

	// 店舗IDを管理する配列(storeIds), 店舗名を管理する配列(storeNames), お気に入り数を管理する配列(favorites) を作成
	// storeIds = append(storeIds, "店舗ID")
	// storeNames = append(storeNames, "店舗名")
	// favorites = append(favorites, time.Now().Format("2006/01/02"))
	// 0~最大値の間で下記の処理を繰り返す(iterator = i)
	// 1. storeIds = append(storeIds, i+1)
	// 2. 店舗ID i+1 が storeInfoMap に存在する場合 => storeNames = append(storeNames, storeInfoMap[i+1])
	//    そうでない場合 => storeNames = append(storeNames, "店舗名なし")
	// 3. 店舗ID i+1 が favoriteMap に存在する場合 => favorites = append(favorites, favoriteMap[i+1])
	//    そうでない場合 => favorites = append(favorites, "0")
	// storeIds, storeNames, favorites を元にSaveDataオブジェクト作成
	// 保存

	// 3. 1. 2. で取得したデータを突合して更新データを生成
	//result := make([]string, len(storeInfo)+1)
	//result = append(result, time.Now().Format("2006/01/02"))
	//for _, e := range storeInfo {
	//	if v, ok := favoriteMap[e.GetId()]; ok {
	//		result = append(result, strconv.Itoa(v))
	//	} else {
	//		result = append(result, "0")
	//	}
	//}

}
