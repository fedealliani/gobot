package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fedealliani/gobot/bot"
	candlestrategy "github.com/fedealliani/gobot/strategies/CandleStrategy"
	percentagestrategy "github.com/fedealliani/gobot/strategies/PercentageStrategy"
)

const CANDLE_STRATEGY = "candleStrategy"
const PERCENTAGE_STRATEGY = "percentageStrategy"

func main() {
	config := GetConfig()
	switch config.StrategyName {
	case CANDLE_STRATEGY:
		candleStrategy, err := candlestrategy.New(config.Variables)
		if err != nil {
			panic(err)
		}
		bot.RunBot(config, candleStrategy)
	case PERCENTAGE_STRATEGY:
		percentageStrategy, err := percentagestrategy.New(config.Variables)
		if err != nil {
			panic(err)
		}
		bot.RunBot(config, percentageStrategy)
	}

}

func GetConfig() bot.Config {
	var err error
	// Default config
	config := bot.Config{
		StrategyName:           "candleStrategy",
		IntervalName:           "1h",
		IntervalInSeconds:      3600,
		Simulator:              true,
		Symbol:                 "BTCUSDT",
		AmountOtherCoinToTrade: 1000,
	}
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("Cannot open config.json. Use default config...")
		return config
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}

	return config
}
