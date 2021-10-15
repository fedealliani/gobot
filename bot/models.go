package bot

type Response [][]interface{}
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
