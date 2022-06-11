package client

type FavoriteClient interface {
	Get() (map[int]int, error)
}
