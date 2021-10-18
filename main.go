package main

import (
	"os"
	"strconv"
	"time"

	"github.com/fedealliani/gobot/bot"
)

func main() {
	var err error
	argsWithoutProg := os.Args[1:]
	interval := "1m"
	simulator := true
	comision := float64(0.00075)
	tradingAmount := int64(15)
	symbol := "BTCDAI"
	intervalToTrade := 31 * time.Second
	amountCandles := int64(3)
	if len(argsWithoutProg) == 7 {
		interval = argsWithoutProg[0]
		simulator, err = strconv.ParseBool(argsWithoutProg[1])
		if err != nil {
			panic(err)
		}
		comision, err = strconv.ParseFloat(argsWithoutProg[2], 64)
		if err != nil {
			panic(err)
		}
		tradingAmount, err = strconv.ParseInt(argsWithoutProg[3], 10, 64)
		if err != nil {
			panic(err)
		}
		symbol = argsWithoutProg[4]
		seconds, err := strconv.ParseInt(argsWithoutProg[5], 10, 64)
		if err != nil {
			panic(err)
		}
		intervalToTrade = time.Duration(seconds) * time.Second
		amountCandles, err = strconv.ParseInt(argsWithoutProg[6], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	bot.RunBot(interval, simulator, comision, tradingAmount, symbol, intervalToTrade, amountCandles)
}
