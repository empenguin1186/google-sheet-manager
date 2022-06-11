package infra

type FavoriteClientImpl struct {
}

func (f FavoriteClientImpl) Get() (map[int]int, error) {
	return map[int]int{1: 1, 100: 1, 101: 2, 102: 3}, nil
}
