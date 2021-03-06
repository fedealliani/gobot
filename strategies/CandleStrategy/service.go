package candlestrategy

import (
	"fmt"

	"github.com/fedealliani/gobot/bot"
)

func (c *CandleStrategy) ShouldBuy(info bot.Info) (bool, float64, error) {
	candlesInfo := GetInfoFromCandles(info.Candles, c.AmountCandles)
	fmt.Println(candlesInfo)
	if candlesInfo.AmountRedCandles == c.AmountCandles {
		tradeAmountWithComission := c.TradingAmount * (1 + info.Comission)
		if tradeAmountWithComission < info.AmountQuoteAssetCoin {
			return true, c.TradingAmount, nil
		} else {
			fmt.Println("Cannot buy because don´t have other coin to buy")
		}
	}

	return false, 0, nil

}

func (c *CandleStrategy) ShouldSell(info bot.Info) (bool, float64, error) {
	candlesInfo := GetInfoFromCandles(info.Candles, c.AmountCandles)
	fmt.Println(candlesInfo)
	if candlesInfo.AmountGreenCandles == c.AmountCandles {
		tradeAmount := c.TradingAmount / info.AvgPrice
		tradeAmountWithComission := tradeAmount * (1 + info.Comission)
		if tradeAmountWithComission < info.AmountBaseAssetCoin {
			return true, tradeAmount, nil
		} else {
			fmt.Println("Cannot sell because don´t have dominant coin to sell")
		}
	}

	return false, 0, nil
}
