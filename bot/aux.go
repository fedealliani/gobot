package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func getCandle(cl []interface{}) Candle {
	open, err := strconv.ParseFloat(cl[1].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	high, err := strconv.ParseFloat(cl[2].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	low, err := strconv.ParseFloat(cl[3].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	close, err := strconv.ParseFloat(cl[4].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}
	volume, err := strconv.ParseFloat(cl[5].(string), 64)
	if err != nil {
		fmt.Errorf("there was an error parsing candles", err)
	}

	candle := Candle{
		OpenTime:  cl[0].(float64),
		Open:      open,
		High:      high,
		Low:       low,
		Close:     close,
		Volume:    volume,
		CloseTime: cl[6].(float64),
	}
	return candle
}
func GetCandlesFromBinance(symbol string, interval string) ([]Candle, error) {
	candles := []Candle{}

	resp, err := http.Get(URL_BINANCE + "?symbol=" + symbol + "&interval=" + interval)
	if err != nil {
		fmt.Printf("there was an error getting candles %v\n", err)
		return candles, err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("there was an invalid status code %d", resp.StatusCode)
		return candles, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Printf("Cannot unmarshal JSON %v\n", err)
		return candles, err
	}

	for _, cl := range result {
		if len(cl) < 7 {
			fmt.Printf("there was an invalid response %v", cl)
			break
		}
		candle := getCandle(cl)
		candles = append(candles, candle)
	}
	return candles, nil
}

func GetValueInUSD(symbol string, value float64) (float64, error) {
	result := 0.0

	resp, err := http.Get("https://cex.io/api/last_price/" + symbol + "/USD")
	if err != nil {
		fmt.Printf("there was an error getting candles %v\n", err)
		return result, err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("there was an invalid status code %d", resp.StatusCode)
		return result, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	var r UsdValue
	err = json.Unmarshal(body, &r)
	if err != nil { // Parse []byte to the go struct pointer
		fmt.Printf("Cannot unmarshal JSON %v\n", err)
		return result, err
	}

	result, err = strconv.ParseFloat(r.LPrice, 64)
	if err != nil {
		fmt.Printf("Cannot converting usd price to float %v\n", err)
		return result, err
	}
	return result * value, nil
}
func Notify(action string) {
	NotifyTelegram(action)
}

func NotifyTelegram(action string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create a telegram service. Ignoring error for demo simplicity.
	token := os.Getenv("TELEGRAM_TOKEN")

	telegramService, _ := telegram.New(token)

	// Passing a telegram chat id as receiver for our messages.
	// Basically where should our message be sent?
	chatIDS := os.Getenv("TELEGRAM_CHAT_ID")

	chatID, err := strconv.ParseInt(chatIDS, 10, 64)
	if err != nil {
		panic(err)
	}
	telegramService.AddReceivers(chatID)

	// Create our notifications distributor.
	notifier := notify.New()

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notifier.UseServices(telegramService)

	// Send a test message.
	_ = notifier.Send(
		context.Background(),
		"Time to making Money",
		action,
	)
}
