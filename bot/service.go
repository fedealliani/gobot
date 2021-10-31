package bot

import (
	"context"
	"fmt"
	"time"
)

func (bot *GoBot) RunBot() {
	config := bot.config
	strategy := bot.strategy
	binanceClient := bot.binanceClient
	fmt.Printf("Configuration: %+v\nStrategy:%+v\n", config, strategy)

	otherCoin := float64(config.AmountOtherCoinToTrade)
	dominantCoin := float64(0)
	firstSymbol := config.Symbol[0:3]
	secondSymbol := config.Symbol[3:]
	for {
		candlesBinance, err := binanceClient.NewKlinesService().Symbol(config.Symbol).
			Interval(config.IntervalName).Do(context.Background())

		if err != nil {
			fmt.Printf("There was an error getting candle from binance. Erorr:%v", err)
			continue
		}
		candles := []Candle{}
		for _, v := range candlesBinance {
			candles = append(candles, getCandle(v))
		}
		info := Info{
			AmountDominantCoin: dominantCoin,
			AmountOtherCoin:    otherCoin,
			Candles:            candles,
		}
		should, amount, err := strategy.ShouldBuy(info)
		if err != nil {
			fmt.Printf("there was an error checking if should buy. Error %v", err)
			continue
		}
		lastPrice := candles[len(candles)-1].Close
		if should {
			if config.Simulator {

				dominantCoin = dominantCoin + amount
				otherCoin = otherCoin - amount*lastPrice
				notificationString := fmt.Sprintf("Buy %.8f %s", amount, firstSymbol)
				fmt.Println(notificationString)
				Notify(notificationString)
			} else {
				// TODO: Buy in binance
			}
		}
		should, amount, err = strategy.ShouldSell(info)
		if err != nil {
			fmt.Printf("there was an error checking if should buy. Error %v", err)
			continue
		}
		if should {
			if config.Simulator {
				dominantCoin = dominantCoin - amount
				otherCoin = otherCoin + amount*lastPrice
				notificationString := fmt.Sprintf("Sell %.8f %s", amount, firstSymbol)
				fmt.Println(notificationString)
				Notify(notificationString)
			} else {
				// TODO: Buy in binance
			}
		}
		firstAmountInUSD, err := GetValueInUSD(firstSymbol, info.AmountDominantCoin)
		if err != nil {
			fmt.Printf("there was an error getting usd value. Error %v", err)
			continue
		}

		secondAmountInUSD, err := GetValueInUSD(secondSymbol, info.AmountOtherCoin)
		if err != nil {
			fmt.Printf("there was an error getting usd value. Error %v", err)
			continue
		}
		totalUsd := firstAmountInUSD + secondAmountInUSD
		usdValueString := fmt.Sprintf("Having %.8f %s and %.8f %s . Representing %.8f USD", info.AmountDominantCoin, firstSymbol, info.AmountOtherCoin, secondSymbol, totalUsd)
		fmt.Println(usdValueString)
		time.Sleep(time.Duration(config.IntervalInSeconds * int64(time.Second)))
	}
}
