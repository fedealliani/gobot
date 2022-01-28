package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/fedealliani/gobot/bot"
)

func (b *BinanceClient) GetAvgPrice(symbol string) (float64, error) {
	priceFloat := float64(0)
	avgPrice, err := b.AvgPrice(symbol)
	if err != nil {
		return priceFloat, err
	}

	if priceFloat, err = strconv.ParseFloat(avgPrice.Price, 64); err != nil {
		return priceFloat, err
	}
	return priceFloat, nil
}
func (b *BinanceClient) GetSymbols(symbol string) (string, string, error) {
	exchangeInfo, err := b.ExchangeInfo(symbol)
	if err != nil {
		return "", "", err
	}
	return exchangeInfo.Symbols[0].BaseAsset, exchangeInfo.Symbols[0].QuoteAsset, nil
}

func (b *BinanceClient) GetCandles(symbol string, interval string) ([]bot.Candle, error) {
	candlesBot := []bot.Candle{}
	candleSticks, err := b.CandleStickData(symbol, interval)
	if err != nil {
		return candlesBot, err
	}
	for _, candleStick := range candleSticks {
		candlesBot = append(candlesBot, getCandle(candleStick))
	}
	return candlesBot, nil
}
func (b *BinanceClient) ExchangeInfo(symbol string) (ExchangeInfo, error) {
	response := ExchangeInfo{}
	queryString := "?symbol=" + symbol
	url := BASE_URL + EXCHANGE_INFO_URL + queryString
	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != http.StatusOK {
		return response, errors.New(fmt.Sprintf("invalid status code %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (b *BinanceClient) CandleStickData(symbol string, interval string) ([]CandleStick, error) {

	response := []CandleStick{}
	queryString := "?symbol=" + symbol + "&" + "interval=" + interval
	resp, err := http.Get(BASE_URL + CANDLESTICK_URL + queryString)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != http.StatusOK {
		return response, errors.New("invalid status code")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	var candlesInterfaces [][]interface{}
	err = json.Unmarshal(body, &candlesInterfaces)
	if err != nil {
		return response, err
	}

	for _, v := range candlesInterfaces {
		candleStick := CandleStick{
			OpenTime:                  v[0].(float64),
			Open:                      v[1].(string),
			High:                      v[2].(string),
			Low:                       v[3].(string),
			Close:                     v[4].(string),
			Volume:                    v[5].(string),
			CloseTime:                 v[6].(float64),
			QuoteAssetVolume:          v[7].(string),
			NumberOfTrades:            v[8].(float64),
			TakerBusyBaseAssetVolume:  v[9].(string),
			TakerBusyQuoteAssetVolume: v[10].(string),
			Ignore:                    v[11].(string),
		}
		response = append(response, candleStick)
	}
	return response, nil
}

func (b *BinanceClient) AvgPrice(symbol string) (AvgPrice, error) {

	response := AvgPrice{}
	queryString := "?symbol=" + symbol
	resp, err := http.Get(BASE_URL + AVGPRICE_URL + queryString)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != http.StatusOK {
		return response, errors.New("invalid status code")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
