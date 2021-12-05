package binance

import (
	"fmt"
	"strconv"

	"github.com/fedealliani/gobot/bot"
)

func getCandle(cl CandleStick) bot.Candle {
	open, err := strconv.ParseFloat(cl.Open, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	high, err := strconv.ParseFloat(cl.High, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	low, err := strconv.ParseFloat(cl.Low, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	close, err := strconv.ParseFloat(cl.Close, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	volume, err := strconv.ParseFloat(cl.Volume, 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}

	candle := bot.Candle{
		OpenTime:  float64(cl.OpenTime),
		Open:      open,
		High:      high,
		Low:       low,
		Close:     close,
		Volume:    volume,
		CloseTime: float64(cl.CloseTime),
	}
	return candle
}
