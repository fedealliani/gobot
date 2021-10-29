package percentagestrategy

import (
	"github.com/fedealliani/gobot/bot"
)

func (c *PercentageStrategy) ShouldBuy(info bot.Info) (bool, float64, error) {
	lastPrice := info.Candles[len(info.Candles)-1].Close
	if !c.Init {
		c.LastPriceTrade = lastPrice
		c.Init = true
		c.BuyMode = true
	}
	if c.BuyMode {
		percentage := 100.0 - (c.LastPriceTrade/lastPrice)*100.0
		if percentage < -c.Percentage {
			c.LastPriceTrade = lastPrice
			c.BuyMode = false
			amountToBuy := info.AmountOtherCoin / lastPrice
			return true, amountToBuy, nil
		}
	}

	return false, 0, nil

}

func (c *PercentageStrategy) ShouldSell(info bot.Info) (bool, float64, error) {
	lastPrice := info.Candles[len(info.Candles)-1].Close
	if !c.Init {
		return false, 0, nil
	}
	if !c.BuyMode {
		percentage := 100.0 - (c.LastPriceTrade/lastPrice)*100.0
		if percentage > c.Percentage {
			c.BuyMode = true
			amountSell := info.AmountDominantCoin * lastPrice
			return true, amountSell, nil
		}

	}

	return false, 0, nil

}
