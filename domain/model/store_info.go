package model

// StoreInfo is the model
type StoreInfo struct {
	// ID of the store.
	id int `json:"id"`
	// Name of the store.
	name string `json:"age"`
}

func (s *StoreInfo) GetId() int {
	return s.id
}

func (s *StoreInfo) GetName() string {
	return s.name
}
