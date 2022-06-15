package model

import "fmt"

type SaveData struct {
	storeIds   []string
	storeNames []string
	favorites  []string
}

func NewSaveData(storeIds []string, storeNames []string, favorites []string) (SaveData, error) {
	if len(storeIds) == len(storeNames) && len(storeNames) == len(favorites) {
		return SaveData{}, fmt.Errorf("each data array must have each length")
	}
	return SaveData{storeIds: storeIds, storeNames: storeNames, favorites: favorites}, nil
}

func (s *SaveData) GetStoreIds() []string {
	return s.storeIds
}

func (s *SaveData) GetStoreNames() []string {
	return s.storeNames
}

func (s *SaveData) GetFavorites() []string {
	return s.favorites
}
