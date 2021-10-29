package candlestrategy

import "github.com/mitchellh/mapstructure"

type CandleStrategy struct {
	AmountCandles int64
	TradingAmount float64
}

func New(variables map[string]interface{}) (*CandleStrategy, error) {
	candleStrategy := CandleStrategy{}
	err := mapstructure.Decode(variables, &candleStrategy)
	return &candleStrategy, err
}
