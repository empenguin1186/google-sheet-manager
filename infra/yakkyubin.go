package infra

import (
	"encoding/json"
	"fmt"
	"github.com/google/martian/log"
	"google-sheet-sample/domain/model"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type YakkyubinClient struct {
	client *http.Client
	config *YakkyubinConfig
}

func NewYakkyubinClient(config *YakkyubinConfig) *YakkyubinClient {
	return &YakkyubinClient{client: &http.Client{}, config: config}
}

func (y *YakkyubinClient) Get() ([]*model.StoreInfo, error) {
	// リクエスト設定
	u := url.URL{}
	u.Scheme = y.config.Scheme
	u.Host = y.config.Host
	u.Path = y.config.Path
	q := u.Query()
	q.Set("chain_id", strconv.Itoa(y.config.ChainId))
	u.RawQuery = q.Encode()
	fmt.Printf("url: %v", u)

	request := http.Request{}
	request.URL = &u
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

	var storeInfo []*model.StoreInfo
	err = json.Unmarshal(body, &storeInfo)
	if err != nil {
		log.Errorf("failed to struct StoreInfo Object: %v", err)
		return []*model.StoreInfo{}, err
	}

	return storeInfo, nil
}
