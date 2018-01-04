package models

// Coin represents a specific CryptoCurrency
type Coin struct {
	ID                  string `json:"Id"`
	URL                 string `json:"Url"`
	ImageURL            string `json:"ImageUrl"`
	Name                string `json:"Name"`
	Symbol              string `json:"Symbol"`
	CoinName            string `json:"CoinName"`
	FullName            string `json:"FullName"`
	Algorithm           string `json:"Algorithm"`
	ProofType           string `json:"ProofType"`
	FullyPremined       string `json:"FullyPremined"`
	TotalCoinSupply     string `json:"TotalCoinSupply"`
	PreMinedValue       string `json:"PreMinedValue"`
	TotalCoinsFreeFloat string `json:"TotalCoinsFreeFloat"`
	SortOrder           string `json:"SortOrder"`
	Sponsored           bool   `json:"Sponsored"`
}

// CoinMarketCapCoin represents a coin from the Coin Market Cap website
type CoinMarketCapCoin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}
