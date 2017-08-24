package datamodel

type MarketSummary struct {
	High    float64
	Low     float64
	Volume  float64
	Last    float64
	Bid     float64
	Ask     float64
	PrevDay float64
}

type Ticker struct {
	Bid  float64 `json:"Bid"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}
