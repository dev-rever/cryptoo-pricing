package repositories

import (
	"os"

	"github.com/dev-rever/cryptoo-pricing/model"
	_ "github.com/dev-rever/cryptoo-pricing/model"
	"github.com/go-resty/resty/v2"
)

const (
	searchQueriesUrl = "https://api.coingecko.com/api/v3/search"
)

type CryptoRepo struct {
	client *resty.Client
}

func ProvideCryptoRepo() *CryptoRepo {
	return &CryptoRepo{
		client: resty.New(),
	}
}

func restfulGet[T any](c *CryptoRepo, url string, queryParams map[string]string) (*T, error) {
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
	if result, err = restfulGet[model.CryptoSearchQueries](c, searchQueriesUrl, params); err != nil {
		return nil, err
	}
	return
}
