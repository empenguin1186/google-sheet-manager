package model

import "fmt"

type SaveData struct {
	storeIds   []int
	storeNames []string
	favorites  []int
}

func NewSaveData(storeIds []int, storeNames []string, favorites []int) (SaveData, error) {
	if len(storeIds) == len(storeNames) && len(storeNames) == len(favorites) {
		return SaveData{storeIds: storeIds, storeNames: storeNames, favorites: favorites}, nil
	}
	return SaveData{}, fmt.Errorf("each data array must have each length")
}

func (s *SaveData) GetStoreIds() []int {
	return s.storeIds
}

func (s *SaveData) GetStoreNames() []string {
	return s.storeNames
}

func (s *SaveData) GetFavorites() []int {
	return s.favorites
}
