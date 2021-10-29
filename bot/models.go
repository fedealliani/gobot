package bot

type Response [][]interface{}
type Info struct {
	AmountDominantCoin float64
	AmountOtherCoin    float64
	Candles            []Candle
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
	StrategyName           string
	IntervalName           string
	IntervalInSeconds      int64
	Simulator              bool
	Symbol                 string
	AmountOtherCoinToTrade float64
	Variables              map[string]interface{}
}
type UsdValue struct {
	LPrice string
}
