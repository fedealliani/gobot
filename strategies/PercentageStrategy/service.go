package percentagestrategy

import (
	"github.com/fedealliani/gobot/bot"
)

func (c *PercentageStrategy) ShouldBuy(info bot.Info) (bool, float64, error) {
	// Check if already define a target
	if !c.Init {
		// Define a init target to buy.
		// The init target is the result of the current price - the configured percentage
		c.Target = (info.AvgPrice * (1 - c.Percentage/100))
		c.Init = true
		c.BuyMode = true
	}

	// Check if bot is in buy mode
	if c.BuyMode {
		// Check if the price is lower than target to buy
		if info.AvgPrice < c.Target {
			// Calculate new target.
			// The sell target is the result of the current price + the configured percentage + two comissions (buy and sell comission)
			c.Target = info.AvgPrice*(1+c.Percentage/100) + 2*info.AmountQuoteAssetCoin*info.Comission
			c.BuyMode = false
			amountToBuy := info.AmountQuoteAssetCoin
			return true, amountToBuy, nil
		}
	}

	return false, 0, nil

}

func (c *PercentageStrategy) ShouldSell(info bot.Info) (bool, float64, error) {

	if !c.BuyMode {
		if info.AvgPrice > c.Target {
			c.BuyMode = true
			c.Target = info.AvgPrice * (1 - c.Percentage/100)
			amountSell := info.AmountBaseAssetCoin
			return true, amountSell, nil
		}

	}

	return false, 0, nil

}
