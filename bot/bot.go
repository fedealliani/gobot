package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func RunBot(config Config) {
	fmt.Printf("Interval %s , simulator %v , comision %v , tradingAmount %v , symbol %s , intervalToTrade %v\n", config.Interval, config.Simulator, config.Comision, config.TradingAmount, config.Symbol, config.IntervalToTrade)
	dai := float64(1000)
	btc := float64(0)

	for {
		resp, err := http.Get(URL_BINANCE + "?symbol=" + config.Symbol + "&interval=" + config.Interval)
		if err != nil {
			fmt.Printf("there was an error getting candles %v\n", err)
		}
		if resp.StatusCode != 200 {
			fmt.Printf("there was an invalid status code %d", resp.StatusCode)
			break

		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body) // response body is []byte

		var result Response
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Printf("Cannot unmarshal JSON %v\n", err)
			break
		}
		candles := []Candle{}
		for _, cl := range result {
			if len(cl) < 7 {
				fmt.Printf("there was an invalid response %v", cl)
				break
			}
			candle := getCandle(cl)
			candles = append(candles, candle)
		}
		candles = candles[len(candles)-int(config.AmountCandles):]
		red := 0
		for _, c := range candles {
			if c.Open > c.Close {
				red++
			}
		}
		if red == int(config.AmountCandles) {
			if config.Simulator {
				if float64(config.TradingAmount) < dai {

					dai -= float64(config.TradingAmount)
					transactionComision := float64(config.TradingAmount) * config.Comision
					dai -= transactionComision
					btcPrice := candles[len(candles)-1].Close
					btcToBuy := float64(config.TradingAmount) / btcPrice
					btc += btcToBuy
					fmt.Printf("BUY %v BTC WITH %v DAI AND %v COMISION IN INTERVAL %s\n", btcToBuy, config.TradingAmount, transactionComision, config.Interval)
					fmt.Printf("DAI:%v BTC:%v TOTALUSD:%v\n", dai, btc, (dai + btc*btcPrice))

				} else {
					fmt.Println("CANNOT BUY, BECAUSE NOT HAVE DAI TO BUY")
				}
			} else {
				// TO DO: Buy in binance
			}
		}
		if red == 0 {
			if config.Simulator {
				btcPrice := candles[len(candles)-1].Close
				btcToSell := float64(config.TradingAmount) / btcPrice
				if btcToSell <= btc {
					btc -= btcToSell
					transactionComision := float64(config.TradingAmount) * config.Comision
					dai -= transactionComision
					dai += float64(config.TradingAmount)
					fmt.Printf("SELL %v BTC WITH %v DAI AND %v COMISION IN INTERVAL %s\n", btcToSell, config.TradingAmount, transactionComision, config.Interval)
					fmt.Printf("DAI:%v BTC:%v TOTALUSD:%v\n", dai, btc, (dai + btc*btcPrice))
				} else {
					fmt.Println("CANNOT SELL, BECAUSE NOT HAVE BTC TO SELL")
				}
			} else {
				// TO DO: SELL in binance
			}
		}

		time.Sleep(config.IntervalToTrade)
	}
}
