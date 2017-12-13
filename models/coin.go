package models

type coin struct {
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
