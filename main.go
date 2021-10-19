package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/fedealliani/gobot/bot"
)

func main() {
	config := GetConfig()
	bot.RunBot(config)
}

func GetConfig() bot.Config {
	var err error
	// Default config
	config := bot.Config{
		Interval:        "1m",
		Simulator:       true,
		Comision:        float64(0.00075),
		TradingAmount:   int64(15),
		Symbol:          "BTCDAI",
		IntervalToTrade: 31 * time.Second,
		AmountCandles:   int64(3),
	}
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 7 {
		config.Interval = argsWithoutProg[0]
		config.Simulator, err = strconv.ParseBool(argsWithoutProg[1])
		if err != nil {
			panic(err)
		}

		config.Comision, err = strconv.ParseFloat(argsWithoutProg[2], 64)
		if err != nil {
			panic(err)
		}
		config.TradingAmount, err = strconv.ParseInt(argsWithoutProg[3], 10, 64)
		if err != nil {
			panic(err)
		}
		config.Symbol = argsWithoutProg[4]
		seconds, err := strconv.ParseInt(argsWithoutProg[5], 10, 64)
		if err != nil {
			panic(err)
		}
		config.IntervalToTrade = time.Duration(seconds) * time.Second
		config.AmountCandles, err = strconv.ParseInt(argsWithoutProg[6], 10, 64)
		if err != nil {
			panic(err)
		}
	} else {
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
		config.IntervalToTrade = config.IntervalToTrade * time.Second
	}
	return config
}
