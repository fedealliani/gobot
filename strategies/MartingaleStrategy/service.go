package martingalestrategy

import (
	"fmt"

	"github.com/fedealliani/gobot/bot"
)

func (c *MartingaleStrategy) ShouldBuy(info bot.Info) (bool, float64, error) {
	if !c.Init {
		if checkIfHaveMoneyToBuy(0, c.AmountFirstTrade, info) {
			purchase := c.AmountFirstTrade / info.LastPrice
			bet := purchase * info.LastPrice
			avg := (bet * (1 + info.Comission) / purchase) + (bet * info.Comission)
			c.Purchases = []float64{purchase}
			c.Bets = []float64{bet}
			c.Averages = []float64{avg}

			c.Init = true
			c.BuyMode = true

			return true, c.Bets[c.Iteration], nil
		} else {
			return false, 0, nil
		}
	} else {

		if c.BuyMode {
			if c.Averages[c.Iteration] >= info.LastPrice {
				if checkIfHaveMoneyToBuy(c.Bets[c.Iteration], c.AmountFirstTrade, info) {

					purchase := 2 * c.Bets[c.Iteration] / info.LastPrice
					bet := purchase * info.LastPrice

					newAveragesNuemerador := float64(0)
					newAveragesDenominador := float64(0)
					totalBets := float64(0)
					for i, _ := range c.Purchases {
						newAveragesNuemerador += c.Bets[i] * (1 + info.Comission)
						newAveragesDenominador += c.Purchases[i]
						totalBets += c.Bets[i]
					}
					totalBets += bet
					newAveragesNuemerador += bet * (1 + info.Comission)
					newAveragesDenominador += purchase
					newAveragesNuemerador += totalBets * info.Comission
					avg := newAveragesNuemerador / newAveragesDenominador

					c.Purchases = append(c.Purchases, purchase)
					c.Bets = append(c.Bets, bet)
					c.Averages = append(c.Averages, avg)
					c.Iteration = c.Iteration + 1

					return true, c.Bets[c.Iteration], nil
				} else {
					fmt.Println("no tengo plata para hacer trading")
				}
			} else {
				c.BuyMode = false
				c.Iteration = 0
				return false, 0, nil
			}
		}

	}

	return false, 0, nil

}

func (c *MartingaleStrategy) ShouldSell(info bot.Info) (bool, float64, error) {
	if !c.BuyMode && c.Init {
		c.BuyMode = true
		c.Init = false
		return true, info.AmountDominantCoin, nil
	}

	return false, 0, nil

}
