package bot

type Exchange interface {
	GetSymbols(pair string) (string, string, error)
	GetCandles(symbol string, interval string) ([]Candle, error)
	GetAvgPrice(symbol string) (float64, error)
}
