package binance

type ErrorResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type ExchangeInfo struct {
	Timezone   string    `json:"timezone"`
	ServerTime int64     `json:"serverTime"`
	Symbols    []Symbols `json:"symbols"`
}
type Symbols struct {
	Symbol                   string   `json:"symbol"`
	Status                   string   `json:"status"`
	BaseAsset                string   `json:"baseAsset"`
	BaseAssetPrecision       int      `json:"baseAssetPrecision"`
	QuoteAsset               string   `json:"quoteAsset"`
	QuotePrecision           int      `json:"quotePrecision"`
	QuoteAssetPrecision      int      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision  int      `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision int      `json:"quoteCommissionPrecision"`
	OrderTypes               []string `json:"orderTypes"`
}

type CandleStick struct {
	OpenTime                  float64
	Open                      string
	High                      string
	Low                       string
	Close                     string
	Volume                    string
	CloseTime                 float64
	QuoteAssetVolume          string
	NumberOfTrades            float64
	TakerBusyBaseAssetVolume  string
	TakerBusyQuoteAssetVolume string
	Ignore                    string
}

type AvgPrice struct {
	Mins  int64
	Price string
}
