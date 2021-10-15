package bot

import (
	"fmt"
	"strconv"
)

func getCandle(cl []interface{}) Candle {
	open, err := strconv.ParseFloat(cl[1].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	high, err := strconv.ParseFloat(cl[2].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	low, err := strconv.ParseFloat(cl[3].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	close, err := strconv.ParseFloat(cl[4].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	volume, err := strconv.ParseFloat(cl[5].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}

	candle := Candle{
		OpenTime:  cl[0].(float64),
		Open:      open,
		High:      high,
		Low:       low,
		Close:     close,
		Volume:    volume,
		CloseTime: cl[6].(float64),
	}
	return candle
}
