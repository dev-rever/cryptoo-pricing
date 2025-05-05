package model

type CryptoSearchQueries struct {
	Coins      []Coin        `json:"coins,omitempty"`
	Exchanges  []Exchange    `json:"exchanges,omitempty"`
	Icos       []interface{} `json:"icos,omitempty"`
	Categories []Category    `json:"categories,omitempty"`
	Nfts       []Nft         `json:"nfts,omitempty"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Coin struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	APISymbol     string `json:"api_symbol"`
	Symbol        string `json:"symbol"`
	MarketCapRank int64  `json:"-"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}

type Exchange struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MarketType string `json:"market_type"`
	Thumb      string `json:"thumb"`
	Large      string `json:"large"`
}

type Nft struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Thumb  string `json:"thumb"`
}
