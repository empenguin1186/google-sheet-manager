package client

type StoreClient interface {
	Get() (map[int]string, int, error)
}
