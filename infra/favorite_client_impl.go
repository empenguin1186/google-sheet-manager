package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type FavoriteClientImpl struct {
	client  *http.Client
	idToken string
	config  *ApiConfig
}

func NewFavoriteClientImpl(idToken string, config *ApiConfig) *FavoriteClientImpl {
	return &FavoriteClientImpl{client: &http.Client{}, idToken: idToken, config: config}
}

func (f FavoriteClientImpl) Get() (map[int]int, error) {
	u, err := url.Parse(f.config.Url)
	if err != nil {
		log.Printf("failed to parse url: %v", err)
		return map[int]int{}, err
	}

	request := http.Request{}
	request.URL = u
	request.Method = "GET"

	header := http.Header{}
	header.Add("Authorization", "Bearer "+f.idToken)
	request.Header = header

	response, err := f.client.Do(&request)
	if err != nil {
		log.Printf("failed to request to favorite api: %v", err)
		return map[int]int{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("status code 200 not returned to favorite api: %v", err)
		return map[int]int{}, fmt.Errorf("status code 200 not returned to favorite api")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return map[int]int{}, err
	}

	var favoriteResponse FavoriteResponse
	err = json.Unmarshal(body, &favoriteResponse)
	if err != nil {
		log.Printf("failed to struct FavoriteInfo Object: %v", err)
		return map[int]int{}, err
	}

	result := map[int]int{}
	for _, e := range favoriteResponse.Results {
		result[e.StoreId] = e.Favorite
	}

	return result, nil
}

type FavoriteInfo struct {
	StoreId  int `json:"storeId"`
	Favorite int `json:"favorite"`
}

type FavoriteResponse struct {
	Results []*FavoriteInfo `json:"results"`
}
