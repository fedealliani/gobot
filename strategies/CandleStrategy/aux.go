package candlestrategy

import "github.com/fedealliani/gobot/bot"

func GetInfoFromCandles(candles []bot.Candle, n int64) CandleStrategyModel {
	candles = candles[len(candles)-int(n):]
	red := 0
	green := 0
	draw := 0
	for _, c := range candles {
		if c.Open > c.Close {
			red++
		} else if c.Open == c.Close {
			draw++
		} else {
			green++
		}
	}
	return CandleStrategyModel{
		AmountRedCandles:   int64(red),
		AmountGreenCandles: int64(green),
		AmountDrawCandles:  int64(draw),
	}
}
