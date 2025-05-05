package repositories

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/dev-rever/cryptoo-pricing/model"
	_ "github.com/dev-rever/cryptoo-pricing/model"
	"github.com/go-resty/resty/v2"
)

const (
	searchQueriesUrl       = "https://api.coingecko.com/api/v3/search"
	supportVsCurrenciesUrl = "https://api.coingecko.com/api/v3/simple/supported_vs_currencies"
	coinPriceByIDsUrl      = "https://api.coingecko.com/api/v3/simple/price"
)

type CryptoRepo struct {
	client *resty.Client
}

func ProvideCryptoRepo() *CryptoRepo {
	return &CryptoRepo{
		client: resty.New(),
	}
}

func restfulGet(c *CryptoRepo, url string, queryParams map[string]string) (response *resty.Response, err error) {
	req := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("x-cg-demo-api-key", os.Getenv("COINGECKO_API_KEY"))

	if queryParams != nil {
		req.SetQueryParams(queryParams)
	}

	response, err = req.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func restfulGetWithType[T any](c *CryptoRepo, url string, queryParams map[string]string) (*T, error) {
	var result T
	req := c.client.R().
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("x-cg-demo-api-key", os.Getenv("COINGECKO_API_KEY"))

	if queryParams != nil {
		req.SetQueryParams(queryParams)
	}

	_, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func restfulPost[T any](c *CryptoRepo, url string, queryParams map[string]string) (*T, error) {
	var result T
	req := c.client.R().
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("x-cg-demo-api-key", os.Getenv("COINGECKO_API_KEY"))

	if queryParams != nil {
		req.SetQueryParams(queryParams)
	}

	_, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *CryptoRepo) SearchQueries(coin string) (result *model.CryptoSearchQueries, err error) {
	var params = map[string]string{"query": coin}
	if result, err = restfulGetWithType[model.CryptoSearchQueries](c, searchQueriesUrl, params); err != nil {
		return nil, err
	}
	return
}

func (c *CryptoRepo) SupportedVsCurrencies() (result *[]string, err error) {
	if result, err = restfulGetWithType[[]string](c, supportVsCurrenciesUrl, map[string]string{}); err != nil {
		return nil, err
	}
	return
}

func (c *CryptoRepo) CoinPriceByIDs(currencies []string, coins []string) (result map[string]map[string]float32, err error) {
	params := map[string]string{
		"include_market_cap":      "true",
		"include_24hr_vol":        "true",
		"include_24hr_change":     "true",
		"include_last_updated_at": "true",
		"precision":               "full",
		"vs_currencies":           strings.Join(currencies, ","),
		"ids":                     strings.Join(coins, ","),
	}

	response, err := restfulGet(c, coinPriceByIDsUrl, params)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
