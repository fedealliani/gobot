package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func RunBot(interval string, simulator bool, comision float64, tradingAmount int64, symbol string, intervalToTrade time.Duration, amountCandles int64) {
	fmt.Printf("Interval %s , simulator %v , comision %v , tradingAmount %v , symbol %s , intervalToTrade %v", interval, simulator, comision, tradingAmount, symbol, intervalToTrade)
	dai := float64(1000)
	btc := float64(0)

	for {
		resp, err := http.Get(URL_BINANCE + "?symbol=" + symbol + "&interval=" + interval)
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
		candles = candles[len(candles)-int(amountCandles):]
		red := 0
		for _, c := range candles {
			if c.Open > c.Close {
				red++
			}
		}
		if red == int(amountCandles) {
			if simulator {
				if float64(tradingAmount) < dai {

					dai -= float64(tradingAmount)
					transactionComision := float64(tradingAmount) * comision
					dai -= transactionComision
					btcPrice := candles[len(candles)-1].Close
					btcToBuy := float64(tradingAmount) / btcPrice
					btc += btcToBuy
					fmt.Printf("BUY %v BTC WITH %v DAI AND %v COMISION IN INTERVAL %s\n", btcToBuy, tradingAmount, transactionComision, interval)
					fmt.Printf("DAI:%v BTC:%v TOTALUSD:%v\n", dai, btc, (dai + btc*btcPrice))

				} else {
					fmt.Println("CANNOT BUY, BECAUSE NOT HAVE DAI TO BUY")
				}
			} else {
				// TO DO: Buy in binance
			}
		}
		if red == 0 {
			if simulator {
				btcPrice := candles[len(candles)-1].Close
				btcToSell := float64(tradingAmount) / btcPrice
				if btcToSell <= btc {
					btc -= btcToSell
					transactionComision := float64(tradingAmount) * comision
					dai -= transactionComision
					dai += float64(tradingAmount)
					fmt.Printf("SELL %v BTC WITH %v DAI AND %v COMISION IN INTERVAL %s\n", btcToSell, tradingAmount, transactionComision, interval)
					fmt.Printf("DAI:%v BTC:%v TOTALUSD:%v\n", dai, btc, (dai + btc*btcPrice))
				} else {
					fmt.Println("CANNOT SELL, BECAUSE NOT HAVE BTC TO SELL")
				}
			} else {
				// TO DO: SELL in binance
			}
		}

		time.Sleep(intervalToTrade)
	}
}
