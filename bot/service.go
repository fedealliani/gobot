package bot

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func (bot *GoBot) RunBot() {
	config := bot.config
	strategy := bot.strategy
	exchangeClient := bot.exchangeClient
	log := logrus.New()
	log.SetLevel(logrus.Level(config.LoggerLvl))

	log.Infof("Configuration: %+v", config)
	log.Infof("Strategy:%+v", strategy)

	amountQuoteAsset := float64(config.AmountQuoteAssetToTrade)
	dominantCoin := float64(0)
	firstSymbol := config.Symbol[0:3]
	secondSymbol := config.Symbol[3:]
	log.Infof("Start with %.2f %s\n", amountQuoteAsset, secondSymbol)

	for {
		candles, err := exchangeClient.GetCandles(config.Symbol, config.IntervalName)
		if err != nil {
			log.Errorf("There was an error getting candle from binance. Erorr:%v", err)
			continue
		}

		avgPrice, err := exchangeClient.GetAvgPrice(config.Symbol)
		if err != nil {
			log.Errorf("There was an error getting avg price from binance. Erorr:%v", err)
			continue
		}
		info := Info{
			AmountBaseAssetCoin:  dominantCoin,
			AmountQuoteAssetCoin: amountQuoteAsset,
			Candles:              candles,
			Comission:            0.1 / 100,
			AvgPrice:             avgPrice,
		}
		should, amount, err := strategy.ShouldBuy(info)
		if err != nil {
			log.Errorf("there was an error checking if should buy. Error %v", err)
			continue
		}

		if should {
			if config.Simulator {
				comision := amount * info.Comission
				totalBuyAmount := amount - comision
				dominantCoin = dominantCoin + totalBuyAmount/avgPrice
				amountQuoteAsset = amountQuoteAsset - amount
				notificationString := fmt.Sprintf("Buy %.8f %s", amount, secondSymbol)
				log.Debug(notificationString)
				profit := dominantCoin*avgPrice + amountQuoteAsset - config.AmountQuoteAssetToTrade
				usdValueString := fmt.Sprintf("Having %.8f %s and %.8f %s . Profit %.8f %s", dominantCoin, firstSymbol, amountQuoteAsset, secondSymbol, profit, secondSymbol)
				log.Debug(usdValueString)
				//Notify(notificationString)
			} else {
				// TODO: Buy in binance
			}
		}
		should, amount, err = strategy.ShouldSell(info)
		if err != nil {
			log.Errorf("there was an error checking if should buy. Error %v", err)
			continue
		}

		if should {
			if config.Simulator {
				comision := amount * avgPrice * info.Comission
				totalMoneySell := amount*avgPrice - comision
				dominantCoin = dominantCoin - amount
				amountQuoteAsset = amountQuoteAsset + totalMoneySell
				notificationString := fmt.Sprintf("Sell %.8f %s", amount, firstSymbol)
				log.Debug(notificationString)
				profit := dominantCoin*avgPrice + amountQuoteAsset - config.AmountQuoteAssetToTrade
				usdValueString := fmt.Sprintf("Having %.8f %s and %.8f %s . Profit %.8f %s", dominantCoin, firstSymbol, amountQuoteAsset, secondSymbol, profit, secondSymbol)
				log.Debug(usdValueString)
				//Notify(notificationString)
			} else {
				// TODO: Buy in binance
			}
		}

		time.Sleep(time.Duration(config.IntervalInSeconds * int64(time.Second)))
	}
}
