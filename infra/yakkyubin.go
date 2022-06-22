package infra

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type YakkyubinClient struct {
	client *http.Client
	config *ApiConfig
}

func NewYakkyubinClient(config *ApiConfig) *YakkyubinClient {
	return &YakkyubinClient{client: &http.Client{}, config: config}
}

// Get 薬急便APIを実行して店舗IDと店舗名のkey-value形式のデータおよびIDの最大値を取得する
func (y *YakkyubinClient) Get() (map[int]string, int, error) {
	u, err := url.Parse(y.config.Url)
	if err != nil {
		log.Printf("failed to parse url: %v", err)
		return map[int]string{}, 0, err
	}

	request := http.Request{}
	request.URL = u
	request.Method = "GET"

	response, err := y.client.Do(&request)
	if err != nil {
		log.Printf("failed to request to yakkyubin api: %v", err)
		return map[int]string{}, 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("status code 200 not returned to yakkyubin api: %v", err)
		return map[int]string{}, 0, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return map[int]string{}, 0, err
	}

	var yakkyubinResponse YakkyubinResponse
	err = json.Unmarshal(body, &yakkyubinResponse)
	if err != nil {
		log.Printf("failed to struct StoreInfo Object: %v", err)
		return map[int]string{}, 0, err
	}

	result := map[int]string{}
	maximum := 0
	for _, e := range yakkyubinResponse.Pharmacies {
		result[e.Id] = e.Name
		if maximum < e.Id {
			maximum = e.Id
		}
	}

	return result, maximum, nil
}

// StoreInfo is the model
type StoreInfo struct {
	// ID of the store.
	Id int `json:"id"`
	// Name of the store.
	Name string `json:"name"`
}

type YakkyubinResponse struct {
	Pharmacies []*StoreInfo `json:"pharmacies"`
}
