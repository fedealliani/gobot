package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func (bot *GoBot) RunBot() {
	config := bot.config
	strategy := bot.strategy
	binanceClient := bot.binanceClient
	log := logrus.New()
	log.SetLevel(logrus.Level(config.LoggerLvl))

	log.Infof("Configuration: %+v", config)
	log.Infof("Strategy:%+v", strategy)

	otherCoin := float64(config.AmountQuoteAssetToTrade)
	dominantCoin := float64(0)
	firstSymbol := config.Symbol[0:3]
	secondSymbol := config.Symbol[3:]
	log.Infof("Start with %.2f %s\n", otherCoin, secondSymbol)

	for {
		candlesBinance, err := binanceClient.NewKlinesService().Symbol(config.Symbol).
			Interval(config.IntervalName).Do(context.Background())

		if err != nil {
			log.Errorf("There was an error getting candle from binance. Erorr:%v", err)
			continue
		}
		candles := []Candle{}
		for _, v := range candlesBinance {
			candles = append(candles, getCandle(v))
		}
		lastPrice := candles[len(candles)-1].Close
		info := Info{
			AmountDominantCoin: dominantCoin,
			AmountOtherCoin:    otherCoin,
			Candles:            candles,
			Comission:          0.1 / 100,
			LastPrice:          lastPrice,
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
				dominantCoin = dominantCoin + totalBuyAmount/lastPrice
				otherCoin = otherCoin - amount
				notificationString := fmt.Sprintf("Buy %.8f %s", amount, secondSymbol)
				log.Debug(notificationString)
				profit := dominantCoin*lastPrice + otherCoin - config.AmountQuoteAssetToTrade
				usdValueString := fmt.Sprintf("Having %.8f %s and %.8f %s . Profit %.8f %s", dominantCoin, firstSymbol, otherCoin, secondSymbol, profit, secondSymbol)
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
				comision := amount * lastPrice * info.Comission
				totalMoneySell := amount*lastPrice - comision
				dominantCoin = dominantCoin - amount
				otherCoin = otherCoin + totalMoneySell
				notificationString := fmt.Sprintf("Sell %.8f %s", amount, firstSymbol)
				log.Debug(notificationString)
				profit := dominantCoin*lastPrice + otherCoin - config.AmountQuoteAssetToTrade
				usdValueString := fmt.Sprintf("Having %.8f %s and %.8f %s . Profit %.8f %s", dominantCoin, firstSymbol, otherCoin, secondSymbol, profit, secondSymbol)
				log.Debug(usdValueString)
				//Notify(notificationString)
			} else {
				// TODO: Buy in binance
			}
		}

		time.Sleep(time.Duration(config.IntervalInSeconds * int64(time.Second)))
	}
}
