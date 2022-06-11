package client

type DataManager interface {
	Save(data []string) error
}
