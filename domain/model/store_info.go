package model

// StoreInfo is the model
type StoreInfo struct {
	// ID of the store.
	Id int `json:"id"`
	// Name of the store.
	Name string `json:"name"`
}

func (s *StoreInfo) GetId() int {
	return s.Id
}

func (s *StoreInfo) GetName() string {
	return s.Name
}
