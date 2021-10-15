package main

import (
	"time"

	"github.com/fedealliani/gobot/bot"
)

func main() {
	interval := "1m"
	simulator := false
	comision := float64(0.00075)
	tradingAmount := 15
	symbol := "BTCDAI"
	intervalToTrade := 31 * time.Second
	bot.RunBot(interval, simulator, comision, int64(tradingAmount), symbol, intervalToTrade)
}
