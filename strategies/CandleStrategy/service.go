package candlestrategy

import (
	"fmt"

	"github.com/fedealliani/gobot/bot"
)

func (c *CandleStrategy) ShouldBuy(info bot.Info) (bool, float64, error) {
	candlesInfo := GetInfoFromCandles(info.Candles, c.AmountCandles)
	fmt.Println(candlesInfo)
	if candlesInfo.AmountRedCandles == c.AmountCandles {
		otherCoinAmount := c.TradingAmount * info.Candles[len(info.Candles)-1].Close
		if otherCoinAmount < info.AmountOtherCoin {
			return true, c.TradingAmount, nil
		} else {
			fmt.Println("Cannot buy because don´t have other coin to buy")
		}
	}

	return false, 0, nil

}

func (c *CandleStrategy) ShouldSell(info bot.Info) (bool, float64, error) {
	candlesInfo := GetInfoFromCandles(info.Candles, c.AmountCandles)
	if candlesInfo.AmountGreenCandles == c.AmountCandles {

		if c.TradingAmount < info.AmountDominantCoin {
			return true, c.TradingAmount, nil
		} else {
			fmt.Println("Cannot sell because don´t have dominant coin to sell")
		}
	}

	return false, 0, nil
}
