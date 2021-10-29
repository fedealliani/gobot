package percentagestrategy

import "github.com/mitchellh/mapstructure"

type PercentageStrategy struct {
	Percentage     float64
	LastPriceTrade float64
	BuyMode        bool
	Init           bool
}

func New(variables map[string]interface{}) (*PercentageStrategy, error) {
	percentageStrategy := PercentageStrategy{}
	err := mapstructure.Decode(variables, &percentageStrategy)
	if err != nil {
		return nil, err
	}
	return &percentageStrategy, err
}
