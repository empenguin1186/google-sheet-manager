package infra

import (
	"encoding/json"
	"github.com/google/martian/log"
	"google-sheet-sample/domain/model"
	"io"
	"net/http"
	"net/url"
)

type YakkyubinClient struct {
	client *http.Client
	config *YakkyubinConfig
}

func NewYakkyubinClient(config *YakkyubinConfig) *YakkyubinClient {
	return &YakkyubinClient{client: &http.Client{}, config: config}
}

func (y *YakkyubinClient) Get() ([]*model.StoreInfo, error) {
	u, err := url.Parse(y.config.Url)
	if err != nil {
		log.Errorf("failed to parse url: %v", err)
		return []*model.StoreInfo{}, err
	}

	request := http.Request{}
	request.URL = u
	request.Method = "GET"

	response, err := y.client.Do(&request)
	if err != nil {
		log.Errorf("failed to request to yakkyubin api: %v", err)
		return []*model.StoreInfo{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Errorf("status code 200 not returned to yakkyubin api: %v", err)
		return []*model.StoreInfo{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("failed to read response body: %v", err)
		return []*model.StoreInfo{}, err
	}

	var yakkyubinResponse YakkyubinResponse
	err = json.Unmarshal(body, &yakkyubinResponse)
	if err != nil {
		log.Errorf("failed to struct StoreInfo Object: %v", err)
		return []*model.StoreInfo{}, err
	}

	return yakkyubinResponse.Pharmacies, nil
}

type YakkyubinResponse struct {
	Pharmacies []*model.StoreInfo `json:"pharmacies"`
}
