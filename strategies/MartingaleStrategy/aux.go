package martingalestrategy

import "github.com/fedealliani/gobot/bot"

func calculateNewBet(lastBet float64, info bot.Info) float64 {
	return 2 * lastBet * (1 + info.Comission)
}
func checkIfHaveMoneyToBuy(lastBet float64, firstTrade float64, info bot.Info) bool {
	if lastBet == 0 {
		return firstTrade*(1+info.Comission) < info.AmountOtherCoin
	}
	return calculateNewBet(lastBet, info) < info.AmountOtherCoin
}
