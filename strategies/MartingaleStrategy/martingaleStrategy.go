package martingalestrategy

import "github.com/mitchellh/mapstructure"

type MartingaleStrategy struct {
	AmountFirstTrade float64
	BuyMode          bool
	Init             bool
	Purchases        []float64
	Averages         []float64
	Bets             []float64
	Iteration        int64
}

func New(variables map[string]interface{}) (*MartingaleStrategy, error) {
	martingaleStrategy := MartingaleStrategy{}
	err := mapstructure.Decode(variables, &martingaleStrategy)
	if err != nil {
		return nil, err
	}
	return &martingaleStrategy, err
}
