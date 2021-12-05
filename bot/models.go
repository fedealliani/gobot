package bot

type Response [][]interface{}
type Info struct {
	AmountDominantCoin float64
	AmountOtherCoin    float64
	Candles            []Candle
	Comission          float64
	LastPrice          float64
}
type Candle struct {
	OpenTime                  float64
	Open                      float64
	High                      float64
	Low                       float64
	Close                     float64
	Volume                    float64
	CloseTime                 float64
	QuoteAssetVolume          float64
	NumberOfTrades            int64
	TakerBusyBaseAssetVolume  float64
	TakerBusyQuoteAssetVolume float64
	Ignore                    float64
}
type Config struct {
	StrategyName            string                 `json:"strategyName"`
	IntervalName            string                 `json:"intervalName"`
	LoggerLvl               int64                  `json:"loggerLvl"`
	IntervalInSeconds       int64                  `json:"intervalInSeconds"`
	Simulator               bool                   `json:"simulator"`
	Symbol                  string                 `json:"symbol"`
	AmountQuoteAssetToTrade float64                `json:"amountQuoteAssetToTrade"`
	Variables               map[string]interface{} `json:"variables"`
}
type UsdValue struct {
	LPrice string
}
