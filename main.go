package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/fedealliani/gobot/bot"
	candlestrategy "github.com/fedealliani/gobot/strategies/CandleStrategy"
	martingalestrategy "github.com/fedealliani/gobot/strategies/MartingaleStrategy"
	percentagestrategy "github.com/fedealliani/gobot/strategies/PercentageStrategy"
	"github.com/joho/godotenv"
)

const CANDLE_STRATEGY = "candleStrategy"
const PERCENTAGE_STRATEGY = "percentageStrategy"
const MARTINGALE_STRATEGY = "martingaleStrategy"

func main() {
	config := GetConfig()
	binanceClient := GetBinanceClient()

	switch config.StrategyName {
	case CANDLE_STRATEGY:
		candleStrategy, err := candlestrategy.New(config.Variables)
		if err != nil {
			panic(err)
		}
		bot := bot.New(candleStrategy, config, binanceClient)
		bot.RunBot()
	case PERCENTAGE_STRATEGY:
		percentageStrategy, err := percentagestrategy.New(config.Variables)
		if err != nil {
			panic(err)
		}
		bot := bot.New(percentageStrategy, config, binanceClient)
		bot.RunBot()
	case MARTINGALE_STRATEGY:
		martingaleStrategy, err := martingalestrategy.New(config.Variables)
		if err != nil {
			panic(err)
		}
		bot := bot.New(martingaleStrategy, config, binanceClient)
		bot.RunBot()
	}

}

func GetConfig() bot.Config {
	var err error
	// Default config
	config := bot.Config{
		StrategyName:            "candleStrategy",
		IntervalName:            "1h",
		LoggerLvl:               2,
		IntervalInSeconds:       3600,
		Simulator:               true,
		Symbol:                  "BTCUSDT",
		AmountQuoteAssetToTrade: 1000,
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

func GetBinanceClient() *binance.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create a telegram service. Ignoring error for demo simplicity.
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	return binance.NewClient(apiKey, secretKey)
}
